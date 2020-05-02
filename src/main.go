package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/elbuki/ctrl-api/src/handler"
)

func main() {
	port := ":" + conf.APIPort
	api := &handler.API{
		Conf:            conf,
		SavedPassphrase: passphraseHash,
	}

	if err := rpc.Register(api); err != nil {
		log.Fatalf("could not register router: %v", err)
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not initialize server: %v", err)
	}

	defer listener.Close()

	log.Printf("CTRL server started at port %s", port)

	rpc.Accept(listener)
}
