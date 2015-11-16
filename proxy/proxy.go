package proxy

import (
	"net/http"
	"time"

	"github.com/eBay/fabio/config"

	gometrics "github.com/eBay/fabio/_third_party/github.com/rcrowley/go-metrics"
)

// Proxy is a dynamic reverse proxy.
type Proxy struct {
	tr       http.RoundTripper
	cfg      config.Proxy
	requests gometrics.Timer
}

func New(tr http.RoundTripper, cfg config.Proxy) *Proxy {
	return &Proxy{
		tr:       tr,
		cfg:      cfg,
		requests: gometrics.GetOrRegisterTimer("requests", gometrics.DefaultRegistry),
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ShuttingDown() {
		http.Error(w, "shutting down", http.StatusServiceUnavailable)
		return
	}

	t := target(r)
	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := addHeaders(r, p.cfg); err != nil {
		http.Error(w, "cannot parse "+r.RemoteAddr, http.StatusInternalServerError)
		return
	}

	var h http.Handler
	switch {
	case r.Header.Get("Upgrade") == "websocket":
		// to establish WS connections as an unfiltered TCP stream
		// between the client and the server use
		//
		//     h = newRawProxy(t.URL)
		//
		// This may resolve compatilbility issues between the client
		// and the server but it also increases the chance for malicious
		// attacks since fabio no longer acts as a middleman.
		h = newWSProxy(t.URL)
	default:
		h = newHTTPProxy(t.URL, p.tr)
	}

	start := time.Now()
	h.ServeHTTP(w, r)
	p.requests.UpdateSince(start)
	t.Timer.UpdateSince(start)
}
