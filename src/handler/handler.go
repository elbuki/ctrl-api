package handler

import "github.com/elbuki/ctrl-api/src/config"

type api struct {
	token           []byte
	conf            config.Config
	savedPassphrase []byte
}

type Handler struct {
	api *api
}

func (h *Handler) SetPassphrase(passphrase []byte) {
	h.api.savedPassphrase = passphrase
}

func NewHandler(c config.Config) *Handler {
	return &Handler{
		api: &api{conf: c},
	}
}
