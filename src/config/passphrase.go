package config

import (
	"fmt"
	"log"
	"syscall"

	"github.com/elbuki/ctrl-api/src/bcrypt"

	"golang.org/x/crypto/ssh/terminal"
)

func (c *Config) GetPassphrase() ([]byte, error) {
	pass, err := askPassphrase()
	if err != nil {
		log.Fatalln(err)
	}

	passphraseHash, err := generatePassphrase(c.Encryptor, pass)
	if err != nil {
		log.Fatalln(err)
	}

	return passphraseHash, nil
}

func askPassphrase() (pass []byte, err error) {
	fmt.Printf("Enter a passphrase: ")

	// Using terminal library for cross os compatibility
	pass, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return pass, fmt.Errorf("could not receive the passphrase: %v\n", err)
	}

	fmt.Print("\n")

	return
}

func generatePassphrase(
	encrypt bcrypt.Encryptor,
	pass []byte,
) ([]byte, error) {

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
