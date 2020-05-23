package handler

import (
	"github.com/elbuki/ctrl-api/src/config"
)

type API struct {
	token           []byte
	conf            config.Config
	savedPassphrase []byte
}

func NewAPI(c config.Config, p []byte) *API {
	return &API{conf: c, savedPassphrase: p}
}
