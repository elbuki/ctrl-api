package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"

	"github.com/elbuki/ctrl-api/src/handler"
	pb "github.com/elbuki/ctrl-protobuf/src/golang"
	"google.golang.org/grpc"
)

func main() {
	port := ":" + conf.APIPort
	srv := grpc.NewServer()
	h := handler.NewHandler(conf)

	h.SetPassphrase(passphraseHash)
	pb.RegisterMainServiceServer(srv, h)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not initialize listener: %v\n", err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	go func() {
		log.Printf("CTRL server started at port %s", port)

		if err := srv.Serve(l); err != nil {
			log.Fatalf("could not start the server: %v\n", err)
		}
	}()

	<-shutdown

	fmt.Println("setting sudo permissions back to normal")

	cmd := exec.Command("/bin/sh", "-c", "sudo chmod 600 /dev/uinput")
	if err := cmd.Run(); err != nil {
		log.Fatalf("could not change permission from uinput: %v", err)
	}

	srv.GracefulStop()
}
