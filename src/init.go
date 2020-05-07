package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/elbuki/ctrl-api/src/bcrypt"
	"github.com/elbuki/ctrl-api/src/config"
)

var (
	conf           config.Config
	passphraseHash []byte
)

func init() {
	if err := conf.SetFlags(nil); err != nil {
		log.Fatalf("could not set command flags: %v", err)
	}

	fmt.Println("sudo permissions are needed to interact with the keyboard")

	cmd := exec.Command("/bin/sh", "-c", "sudo chmod +0666 /dev/uinput")
	if err := cmd.Run(); err != nil {
		log.Fatalf("could not change permission from uinput: %v", err)
	}

	if err := conf.SetController(); err != nil {
		log.Fatalf("could not set keys controller: %v", err)
	}

	if !conf.UsePassphrase {
		return
	}

	conf.Encryptor = bcrypt.NewEncryptor(conf.HashCost)

	pHash, err := conf.GetPassphrase(nil)
	if err != nil {
		log.Fatalf("could not get passphrase: %v", err)
	}

	passphraseHash = pHash
}
