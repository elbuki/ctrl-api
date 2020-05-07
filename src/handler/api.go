package handler

import (
	"github.com/elbuki/ctrl-api/src/config"
)

type API struct {
	Conf            config.Config
	SavedPassphrase []byte
}
