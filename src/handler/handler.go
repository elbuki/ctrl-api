package handler

import (
	"github.com/elbuki/ctrl-api/src/config"
	pb "github.com/elbuki/ctrl-protobuf/proto"
)

type api struct {
	token           string
	conf            config.Config
	savedPassphrase []byte
}

type Handler struct {
	api *api

	pb.UnimplementedMainServiceServer
}

func (h *Handler) SetPassphrase(passphrase []byte) {
	h.api.savedPassphrase = passphrase
}

func NewHandler(c config.Config, passphrase []byte) *Handler {
	return &Handler{
		api: &api{conf: c, savedPassphrase: passphrase},
	}
}
