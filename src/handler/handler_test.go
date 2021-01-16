package handler

import "testing"

func TestRequestMapping(t *testing.T) {

	var response = requestMapping("/ping")
	if response != "pong" {
		t.Errorf("Request ping, expected response pong, got %s", response)
	}
	response = requestMapping("/pong")
	if response != "ping" {
		t.Errorf("Request pong, expected response ping, got %s", response)
	}
	response = requestMapping("/duck")
	if response != "unknown" {
		t.Errorf("Request duck, expected response unknown, got %s", response)
	}
}
