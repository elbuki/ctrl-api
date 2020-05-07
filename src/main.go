package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"os/signal"

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

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		log.Printf("CTRL server started at port %s", port)

		rpc.Accept(listener)
	}()

	<-shutdown

	fmt.Println("sudo permissions are needed to interact with the keyboard")

	cmd := exec.Command("/bin/sh", "-c", "sudo chmod 600 /dev/uinput")
	if err := cmd.Run(); err != nil {
		log.Fatalf("could not change permission from uinput: %v", err)
	}

	if err := listener.Close(); err != nil {
		log.Fatalf("could not close listener: %v", err)
	}
}
