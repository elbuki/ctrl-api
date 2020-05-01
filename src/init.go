package main

import (
	"log"

	"github.com/elbuki/ctrl-api/src/config"
)

var (
	conf           config.Config
	passphraseHash []byte
)

func init() {
	conf.SetFlags()

	if err := conf.SetController(); err != nil {
		log.Fatalf("could not set keys controller: %v", err)
	}

	if !conf.UsePassphrase {
		return
	}

	conf.SetEncryptor()

	pHash, err := conf.GetPassphrase()
	if err != nil {
		log.Fatalf("could not get passphrase: %v", err)
	}

	passphraseHash = pHash
}
