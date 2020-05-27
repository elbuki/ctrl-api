package handler_test

import (
	"context"
	"testing"

	"github.com/elbuki/ctrl-api/src/config"
	"github.com/elbuki/ctrl-api/src/handler"
	pb "github.com/elbuki/ctrl-protobuf/src/golang"
)

type loginScenario struct {
	conf        config.Config
	req         *pb.LoginRequest
	res         *pb.LoginResponse
	shouldThrow bool
}

var loginTable = []loginScenario{
	loginScenario{
		req: &pb.LoginRequest{Uuid: "1234"},
		res: &pb.LoginResponse{},
	},
	loginScenario{
		conf: config.Config{
			UsePassphrase: true,
		},
		req:         &pb.LoginRequest{Uuid: "1234", Passphrase: []byte{}},
		res:         &pb.LoginResponse{Token: []byte{}},
		shouldThrow: true,
	},
	loginScenario{
		req:         &pb.LoginRequest{Passphrase: []byte{}},
		res:         &pb.LoginResponse{Token: []byte{}},
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
