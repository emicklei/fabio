package gcp

/*
Meta data related to fabio provided by compute instances use the following key,value format:

key: fabio
value: scheme=http&port=8080&path=/
*/

import (
	"log"
	"strings"
	"time"

	"github.com/eBay/fabio/_third_party/golang.org/x/oauth2"
	"github.com/eBay/fabio/_third_party/golang.org/x/oauth2/google"
	"github.com/eBay/fabio/_third_party/google.golang.org/api/compute/v1"
)

const FabioKey = "fabio"

type GoogleCloudPlatform struct {
	SecondsBetweenUpdate int
	Project              string
	Zone                 string
}

// https://cloud.google.com/compute/docs/metadata?hl=en
type metadataService struct {
	routeInstructions chan string
	computeService    *compute.Service
	config            GoogleCloudPlatform
}

// NewMetadataService returns a new metadataService provides a backend implementation
// that periodically queries the metadataservice of the Google Cloud Platform
func NewMetadataService(cfg GoogleCloudPlatform) (*metadataService, error) {
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
		config:            cfg,
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
