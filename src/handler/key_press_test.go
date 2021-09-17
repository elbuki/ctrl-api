package handler_test

import (
	"context"
	"testing"

	"github.com/elbuki/ctrl-api/src/config"
	"github.com/elbuki/ctrl-api/src/control"
	"github.com/elbuki/ctrl-api/src/handler"
	pb "github.com/elbuki/ctrl-protobuf/proto"
)

type keyPressScenario struct {
	req         *pb.KeyPressRequest
	shouldThrow bool
}

var keyPressTable = []keyPressScenario{
	keyPressScenario{
		req:         &pb.KeyPressRequest{Key: 9999},
		shouldThrow: true,
	},
	keyPressScenario{
		req:         &pb.KeyPressRequest{Key: 2, Token: []byte("test")},
		shouldThrow: false,
	},
}

func TestKeyPressEndpoint(t *testing.T) {
	conf := config.Config{
		Controller: &control.Controller{},
	}

	for _, s := range keyPressTable {
		h := handler.NewHandler(conf, nil)
		_, err := h.KeyPress(context.Background(), s.req)
		if !s.shouldThrow && err != nil {
			t.Errorf("could not execute key press handler: %v", err)
		}

		if s.shouldThrow && err == nil {
			t.Error("expected error but it went fine")
		}
	}
}
