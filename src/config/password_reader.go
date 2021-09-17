package config

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

type PasswordReader interface {
	ReadPassword() ([]byte, error)
}

type StdinPasswordReader struct{}

func (StdinPasswordReader) ReadPassword() ([]byte, error) {
	fmt.Printf("Enter a passphrase: ")

	// Using terminal library for cross os compatibility
	pass, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return pass, fmt.Errorf("could not receive the passphrase: %v", err)
	}

	fmt.Print("\n")

	return pass, nil
}
