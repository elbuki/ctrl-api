package handler

import (
	"log"

	"github.com/elbuki/ctrl-api/src/config"
)

type API struct {
	Conf config.Config
}

func (a *API) Test(message string, reply *string) error {
	*reply = "Pong"

	log.Println(message)
	log.Printf("Conf: %v", a.Conf)

	return nil
}
