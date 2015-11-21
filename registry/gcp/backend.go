package gcp

/*
Meta data related to fabio provided by compute instances use the following key,value format:

key: fabio
value: scheme=http&port=8080&path=/
*/

import (
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

const FabioKey = "fabio"

type Config struct {
	SecondsBetweenUpdate int
	Project              string
	Zone                 string
}

// for now read config from environment vars passed in by docker
func readConfig() Config {
	return Config{
		SecondsBetweenUpdate: 30,
		Project:              os.Getenv("GCP_PROJECT"),
		Zone:                 os.Getenv("GCP_ZONE"),
	}
}

// https://cloud.google.com/compute/docs/metadata?hl=en
type metadataService struct {
	routeInstructions chan string
	computeService    *compute.Service
	config            Config
}

func NewMetadataService() (*metadataService, error) {
	client, err := google.DefaultClient(oauth2.NoContext, "https://www.googleapis.com/auth/compute")
	if err != nil {
		return nil, err
	}
	computeService, err := compute.New(client)
	if err != nil {
		return nil, err
	}
	service := &metadataService{
		routeInstructions: make(chan string),
		computeService:    computeService,
		config:            readConfig(),
	}
	go service.poll()
	return service, nil
}

func (m *metadataService) ConfigURL() string {
	return ""
}

// Watch watches the services and manual overrides for changes
// and pushes them if there is a difference.
func (m *metadataService) Watch() chan string {
	return m.routeInstructions
}

// https://cloud.google.com/compute/docs/api-rate-limits
func (m *metadataService) poll() {
	for {
		// ask for all running instances
		instances := compute.NewInstancesService(m.computeService)
		call := instances.List(m.config.Project, m.config.Zone)
		list, err := call.Do()
		if err != nil {
			log.Printf("[ERROR] get instances failed %v", err)
			goto sleep
		}
		// for each instance, fetch its metadata
		for _, each := range list.Items {
			getCall := instances.Get(m.config.Project, m.config.Zone, each.Name)
			instance, err := getCall.Do()
			if err != nil {
				log.Printf("[ERROR] get instance failed %v", err)
				goto sleep
			}
			newInstructions := []string{}
			for _, other := range instance.Metadata.Items {
				// for each fabio spec, add the build instruction
				if FabioKey == other.Key {
					// if a build fails then an empty command is added
					newInstructions = append(newInstructions, buildInstruction(instance, other.Value))
				}
			}
			log.Println(strings.Join(newInstructions, "\n"))
			//m.routeInstructions <- strings.Join(newInstructions, "\n")
		}
	sleep:
		time.Sleep(time.Duration(m.config.SecondsBetweenUpdate) * time.Second)
	}
}
