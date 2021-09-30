package handler_test

import (
	"context"
	"testing"

	"github.com/elbuki/ctrl-api/src/config"
	"github.com/elbuki/ctrl-api/src/handler"
	pb "github.com/elbuki/ctrl-protobuf/proto"
)

type loginScenario struct {
	conf        config.Config
	req         *pb.LoginRequest
	shouldThrow bool
}

var loginTable = []loginScenario{
	{
		req: &pb.LoginRequest{Uuid: "1234"},
	},
	{
		conf: config.Config{
			UsePassphrase: true,
		},
		req:         &pb.LoginRequest{Uuid: "1234", Passphrase: []byte{}},
		shouldThrow: true,
	},
	{
		req:         &pb.LoginRequest{Passphrase: []byte{}},
		shouldThrow: true,
	},
}

func TestLoginEndpoint(t *testing.T) {
	for _, s := range loginTable {
		h := handler.NewHandler(s.conf, nil)
		_, err := h.Login(context.Background(), s.req)
		if !s.shouldThrow && err != nil {
			t.Errorf("could not execute pair handler: %v", err)
		}

		if s.shouldThrow && err == nil {
			t.Error("expected error but it went fine")
		}
	}
}
