package tests

import (
	"testing"

	"github.com/saichler/l8types/go/ifs"
)

func TestNetworkMode(t *testing.T) {
	// Test SetNetworkMode and NetworkMode_Native
	ifs.SetNetworkMode(ifs.NETWORK_NATIVE)
	if !ifs.NetworkMode_Native() {
		t.Error("Expected NetworkMode_Native to return true")
	}
	if ifs.NetworkMode_DOCKER() {
		t.Error("Expected NetworkMode_DOCKER to return false")
	}
	if ifs.NetworkMode_K8s() {
		t.Error("Expected NetworkMode_K8s to return false")
	}

	// Test NETWORK_DOCKER
	ifs.SetNetworkMode(ifs.NETWORK_DOCKER)
	if ifs.NetworkMode_Native() {
		t.Error("Expected NetworkMode_Native to return false")
	}
	if !ifs.NetworkMode_DOCKER() {
		t.Error("Expected NetworkMode_DOCKER to return true")
	}
	if ifs.NetworkMode_K8s() {
		t.Error("Expected NetworkMode_K8s to return false")
	}

	// Test NETWORK_K8s
	ifs.SetNetworkMode(ifs.NETWORK_K8s)
	if ifs.NetworkMode_Native() {
		t.Error("Expected NetworkMode_Native to return false")
	}
	if ifs.NetworkMode_DOCKER() {
		t.Error("Expected NetworkMode_DOCKER to return false")
	}
	if !ifs.NetworkMode_K8s() {
		t.Error("Expected NetworkMode_K8s to return true")
	}

	// Reset to native for other tests
	ifs.SetNetworkMode(ifs.NETWORK_NATIVE)
}

func TestNewServiceLink(t *testing.T) {
	asideName := "service-a"
	zsideName := "service-z"
	asideArea := byte(1)
	zsideArea := byte(2)
	mode := ifs.M_RoundRobin
	interval := 100
	request := true

	link := ifs.NewServiceLink(asideName, zsideName, asideArea, zsideArea, mode, interval, request)

	if link.AsideServiceName != asideName {
		t.Errorf("Expected AsideServiceName %s, got %s", asideName, link.AsideServiceName)
	}
	if link.ZsideServiceName != zsideName {
		t.Errorf("Expected ZsideServiceName %s, got %s", zsideName, link.ZsideServiceName)
	}
	if link.AsideServiceArea != int32(asideArea) {
		t.Errorf("Expected AsideServiceArea %d, got %d", asideArea, link.AsideServiceArea)
	}
	if link.ZsideServiceArea != int32(zsideArea) {
		t.Errorf("Expected ZsideServiceArea %d, got %d", zsideArea, link.ZsideServiceArea)
	}
	if link.Interval != uint32(interval) {
		t.Errorf("Expected Interval %d, got %d", interval, link.Interval)
	}
	if link.Request != request {
		t.Errorf("Expected Request %v, got %v", request, link.Request)
	}
	if link.Mode != int32(mode) {
		t.Errorf("Expected Mode %d, got %d", mode, link.Mode)
	}
}
