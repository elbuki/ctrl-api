package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"github.com/elbuki/ctrl-api/src/bcrypt"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	DEFAULT_API_PORT    = "9516"
	DEFAULT_HASH_COST   = 16
	SET_PASSPHRASE_FLAG = "P"
)

var (
	apiPort        string
	usePassphrase  bool
	hashCost       int
	passphraseHash []byte
	encrypt        bcrypt.Encryptor
)

func init() {
	flag.StringVar(
		&apiPort,
		"port",
		DEFAULT_API_PORT,
		"rpc serving port",
	)

	flag.IntVar(
		&hashCost,
		"cost",
		DEFAULT_HASH_COST,
		"hash salt cost for the passphrase",
	)

	flag.BoolVar(
		&usePassphrase,
		SET_PASSPHRASE_FLAG,
		false,
		"use a passphrase for client connections",
	)

	flag.Parse()

	encrypt = bcrypt.Encryptor{}
	encrypt.SetCost(hashCost)

	if !usePassphrase {
		return
	}

	pass, err := askPassphrase()
	if err != nil {
		log.Fatalln(err)
	}

	passphraseHash, err = setPassphrase(pass)
	if err != nil {
		log.Fatalln(err)
	}
}

func askPassphrase() ([]byte, error) {
	var err error
	var pass []byte

	fmt.Printf("Enter a passphrase: ")
	// Using terminal library for cross os compatibility
	pass, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return pass, fmt.Errorf("could not receive the passphrase: %v\n", err)
	}

	fmt.Print("\n")

	return pass, nil
}

func setPassphrase(pass []byte) ([]byte, error) {
	if len(pass) == 0 {
		fmt.Println("Using default passphrase")
		return nil, nil
	}

	passphraseHash, err := encrypt.Generate(pass)
	if err != nil {
		return pass, fmt.Errorf("could not encrypt passphrase: %v", err)
	}

	return passphraseHash, nil
}
