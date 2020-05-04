package main

import (
	"log"

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

	if err := conf.SetController(); err != nil {
		log.Fatalf("could not set keys controller: %v", err)
	}

	if !conf.UsePassphrase {
		return
	}

	conf.Encryptor = bcrypt.NewEncryptor(conf.HashCost)

	pHash, err := conf.GetPassphrase()
	if err != nil {
		log.Fatalf("could not get passphrase: %v", err)
	}

	passphraseHash = pHash
}
