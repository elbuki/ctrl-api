package handler_test

import (
	"context"
	"testing"

	"github.com/elbuki/ctrl-api/src/config"
	"github.com/elbuki/ctrl-api/src/handler"
	"github.com/elbuki/ctrl-api/src/pb"
)

type scenario struct {
	conf        config.Config
	req         *pb.LoginRequest
	res         *pb.LoginResponse
	shouldThrow bool
}

var table = []scenario{
	scenario{
		req: &pb.LoginRequest{Uuid: "1234"},
		res: &pb.LoginResponse{},
	},
	scenario{
		conf: config.Config{
			UsePassphrase: true,
		},
		req:         &pb.LoginRequest{Uuid: "1234", Passphrase: []byte{}},
		res:         &pb.LoginResponse{Token: []byte{}},
		shouldThrow: true,
	},
	scenario{
		req:         &pb.LoginRequest{Passphrase: []byte{}},
		res:         &pb.LoginResponse{Token: []byte{}},
		shouldThrow: true,
	},
}

func TestLoginEndpoint(t *testing.T) {
	for _, s := range table {
		h := handler.NewLoginHandler(handler.NewAPI(s.conf, nil))
		_, err := h.Login(context.Background(), s.req)
		if !s.shouldThrow && err != nil {
			t.Errorf("could not execute pair handler: %v", err)
		}

		if s.shouldThrow && err == nil {
			t.Error("expected error but it went fine")
		}
	}
}
