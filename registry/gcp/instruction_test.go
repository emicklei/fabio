package gcp

import (
	"testing"

	"github.com/eBay/fabio/_third_party/google.golang.org/api/compute/v1"
)

func TestBuildInstruction(t *testing.T) {
	spec := "scheme=http&path=/here&port=8080"
	i := new(compute.Instance)
	i.Name = "test-node-001"
	i.NetworkInterfaces = []*compute.NetworkInterface{&compute.NetworkInterface{NetworkIP: "10.20.30.40"}}
	i.Tags = &compute.Tags{Items: []string{"tic", "tac"}}
	route := buildInstruction(i, spec)
	if got, want := route, "route add test-node-001 /here http://10.20.30.40:8080/here tags tic,tac"; got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBuildInstructionNoTags(t *testing.T) {
	spec := "scheme=http&path=/here&port=8080"
	i := new(compute.Instance)
	i.Name = "test-node-001"
	i.NetworkInterfaces = []*compute.NetworkInterface{&compute.NetworkInterface{NetworkIP: "10.20.30.40"}}
	route := buildInstruction(i, spec)
	if got, want := route, "route add test-node-001 /here http://10.20.30.40:8080/here"; got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBuildInstructionMissingParameter(t *testing.T) {
	for _, each := range []struct {
		spec string
	}{
		{"path=/here&port=8080"},
		{"scheme=http&port=8080"},
		{"scheme=ws&path=/there"},
		{""},
		{"scheme=&path=&port="},
	} {
		i := new(compute.Instance)
		i.Name = "test-node-001"
		i.NetworkInterfaces = []*compute.NetworkInterface{&compute.NetworkInterface{NetworkIP: "10.20.30.40"}}
		route := buildInstruction(i, each.spec)
		if got, want := route, ""; got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
