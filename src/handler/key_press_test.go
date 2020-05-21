package handler_test

import (
	"context"
	"testing"

	"github.com/elbuki/ctrl-api/src/config"
	"github.com/elbuki/ctrl-api/src/handler"
	pb "github.com/elbuki/ctrl-protobuf/src/golang"
)

type keyPressScenario struct {
	conf        config.Config
	req         *pb.KeyPressRequest
	shouldThrow bool
}

// TODO: Initialize controller on tests
// TODO: Test translator by sending an unknown key
// TODO: Test sending an empty token
// TODO: Test sending a different token
// TODO: Test happy path
var keyPressTable = []keyPressScenario{}

func TestKeyPressEndpoint(t *testing.T) {
	for _, s := range keyPressTable {
		h := handler.NewHandler(handler.NewAPI(s.conf, nil))
		_, err := h.KeyPress(context.Background(), s.req)
		if !s.shouldThrow && err != nil {
			t.Errorf("could not execute key press handler: %v", err)
		}

		if s.shouldThrow && err == nil {
			t.Error("expected error but it went fine")
		}
	}
}
