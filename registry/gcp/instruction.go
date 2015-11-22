package gcp

import (
	"bytes"
	"log"
	"net/url"

	"github.com/eBay/fabio/_third_party/google.golang.org/api/compute/v1"
)

// scheme=http&path=/here&port=80
// route add {instance.name} /{spec.path} {spec.scheme}://{instance.networkIP}/{spec.path} {instance.tags.items}
func buildInstruction(instance *compute.Instance, spec string) string {
	v, err := url.ParseQuery(spec)
	if err != nil {
		log.Printf("[ERROR] instance %s has invalid fabio spec %v", instance.Name, err)
		return ""
	}
	path := v.Get("path")
	if !validate(instance.Name, "path", path) {
		return ""
	}
	scheme := v.Get("scheme")
	if !validate(instance.Name, "scheme", scheme) {
		return ""
	}
	if len(instance.NetworkInterfaces) == 0 {
		log.Printf("[ERROR] decoding fabio route parameters %v", err)
		return ""
	}
	ip := instance.NetworkInterfaces[0].NetworkIP
	if !validate(instance.Name, "NetworkIP", ip) {
		return ""
	}
	port := v.Get("port")
	if !validate(instance.Name, "port", port) {
		return ""
	}
	out := new(bytes.Buffer)
	out.WriteString("route add ")
	out.WriteString(instance.Name)
	out.WriteString(" ")
	out.WriteString(path)
	out.WriteString(" ")
	out.WriteString(scheme)
	out.WriteString("://")
	out.WriteString(ip)
	out.WriteString(":")
	out.WriteString(port)
	out.WriteString(path)
	if instance.Tags != nil {
		tagCount := len(instance.Tags.Items)
		if tagCount > 0 {
			out.WriteString(" tags ")
			for i, each := range instance.Tags.Items {
				if i > 0 {
					out.WriteString(",")
				}
				out.WriteString(each)
			}
		}
	}
	return out.String()
}

func validate(instanceName, param, value string) bool {
	if len(value) == 0 {
		log.Printf("[ERROR] instance %s is missing %s parameter for fabio route", instanceName, param)
		return false
	}
	return true
}
