package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/eBay/fabio/admin"
	"github.com/eBay/fabio/config"
	"github.com/eBay/fabio/metrics"
	"github.com/eBay/fabio/proxy"
	"github.com/eBay/fabio/registry"
	"github.com/eBay/fabio/registry/consul"
	"github.com/eBay/fabio/registry/file"
	"github.com/eBay/fabio/registry/gcp"
	"github.com/eBay/fabio/registry/static"
	"github.com/eBay/fabio/route"
)

// version contains the version number
//
// It is set by build/release.sh for tagged releases
// so that 'go get' just works.
//
// It is also set by the linker when fabio
// is built via the Makefile or the build/docker.sh
// script to ensure the correct version nubmer
var version = "1.1.1"

func main() {
	var filename string
	var v bool
	flag.StringVar(&filename, "cfg", "", "path to config file")
	flag.BoolVar(&v, "v", false, "show version")
	flag.Parse()

	if v {
		fmt.Println(version)
		return
	}
	log.Printf("[INFO] Version %s starting", version)

	cfg, err := config.Load(filename)
	if err != nil {
		log.Fatal("[FATAL] ", err)
	}

	initRuntime(cfg)
	initMetrics(cfg)
	initBackend(cfg)
	go watchBackend()
	startAdmin(cfg)
	startListeners(cfg.Listen, cfg.Proxy.ShutdownWait, newProxy(cfg))
	registry.Default.Deregister()
}

func newProxy(cfg *config.Config) *proxy.Proxy {
	if err := route.SetPickerStrategy(cfg.Proxy.Strategy); err != nil {
		log.Fatal("[FATAL] ", err)
	}
	log.Printf("[INFO] Using routing strategy %q", cfg.Proxy.Strategy)

	tr := &http.Transport{
		ResponseHeaderTimeout: cfg.Proxy.ResponseHeaderTimeout,
		MaxIdleConnsPerHost:   cfg.Proxy.MaxConn,
		Dial: (&net.Dialer{
			Timeout:   cfg.Proxy.DialTimeout,
			KeepAlive: cfg.Proxy.KeepAliveTimeout,
		}).Dial,
	}

	return proxy.New(tr, cfg.Proxy)
}

func startAdmin(cfg *config.Config) {
	log.Printf("[INFO] Admin server listening on %q", cfg.UI.Addr)
	go func() {
		if err := admin.ListenAndServe(cfg.UI, version); err != nil {
			log.Fatal("[FATAL] ui: ", err)
		}
	}()
}

func initMetrics(cfg *config.Config) {
	if err := metrics.Init(cfg.Metrics); err != nil {
		log.Fatal("[FATAL] ", err)
	}
}

func initRuntime(cfg *config.Config) {
	if os.Getenv("GOGC") == "" {
		log.Print("[INFO] Setting GOGC=", cfg.Runtime.GOGC)
		debug.SetGCPercent(cfg.Runtime.GOGC)
	} else {
		log.Print("[INFO] Using GOGC=", os.Getenv("GOGC"), " from env")
	}

	if os.Getenv("GOMAXPROCS") == "" {
		log.Print("[INFO] Setting GOMAXPROCS=", cfg.Runtime.GOMAXPROCS)
		runtime.GOMAXPROCS(cfg.Runtime.GOMAXPROCS)
	} else {
		log.Print("[INFO] Using GOMAXPROCS=", os.Getenv("GOMAXPROCS"), " from env")
	}
}

func initBackend(cfg *config.Config) {
	var err error

	switch cfg.Registry.Backend {
	case "file":
		registry.Default, err = file.NewBackend(cfg.Registry.File.Path)
	case "static":
		registry.Default, err = static.NewBackend(cfg.Registry.Static.Routes)
	case "consul":
		registry.Default, err = consul.NewBackend(&cfg.Registry.Consul)
	case "gcp":
		registry.Default, err = gcp.NewBackend(&cfg.Registry.GoogleCloudPlatform)
	default:
		log.Fatal("[FATAL] Unknown registry backend ", cfg.Registry.Backend)
	}

	if err != nil {
		log.Fatal("[FATAL] Error initializing backend. ", err)
	}
	if err := registry.Default.Register(); err != nil {
		log.Fatal("[FATAL] Error registering backend. ", err)
	}
}

func watchBackend() {
	var (
		last   string
		svccfg string
		mancfg string
	)

	svc := registry.Default.WatchServices()
	man := registry.Default.WatchManual()

	for {
		select {
		case svccfg = <-svc:
		case mancfg = <-man:
		}

		// manual config overrides service config
		// order matters
		next := svccfg + "\n" + mancfg
		if next == last {
			continue
		}

		t, err := route.ParseString(next)
		if err != nil {
			log.Printf("[WARN] %s", err)
			continue
		}
		route.SetTable(t)

		last = next
	}
}
