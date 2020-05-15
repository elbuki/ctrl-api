package handler_test

import (
	"testing"

	"github.com/elbuki/ctrl-api/src/config"
	"github.com/elbuki/ctrl-api/src/handler"
)

type scenario struct {
	conf        config.Config
	req         *handler.LoginRequest
	res         *handler.LoginResponse
	shouldThrow bool
}

var table = []scenario{
	scenario{
		req: &handler.LoginRequest{UUID: "1234"},
		res: &handler.LoginResponse{},
	},
	scenario{
		conf: config.Config{
			UsePassphrase: true,
		},
		req:         &handler.LoginRequest{UUID: "1234", Passphrase: []byte{}},
		res:         &handler.LoginResponse{Token: []byte{}},
		shouldThrow: true,
	},
	scenario{
		req:         &handler.LoginRequest{Passphrase: []byte{}},
		res:         &handler.LoginResponse{Token: []byte{}},
		shouldThrow: true,
	},
}

func TestLoginEndpoint(t *testing.T) {
	api := new(handler.API)

	for _, s := range table {
		api.Conf = s.conf

		err := api.Pair(s.req, s.res)
		if !s.shouldThrow && err != nil {
			t.Errorf("could not execute pair handler: %v", err)
		}

		if s.shouldThrow && err == nil {
			t.Error("expected error but it went fine")
		}
	}
}
