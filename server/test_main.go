package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	output := SetupGrpcServer()

	if output != nil {
		t.Errorf("Server not setup properly %v", output)
	}
}
