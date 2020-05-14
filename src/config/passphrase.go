package config

import (
	"fmt"

	"github.com/elbuki/ctrl-api/src/bcrypt"
)

func (c *Config) GetPassphrase(pr PasswordReader) ([]byte, error) {
	if pr == nil {
		pr = &StdinPasswordReader{}
	}

	pass, err := pr.ReadPassword()
	if err != nil {
		return nil, fmt.Errorf("could not read password: %v", err)
	}

	passphraseHash, err := generatePassphrase(c.Encryptor, pass)
	if err != nil {
		return nil, fmt.Errorf("could not generate pass: %v", err)
	}

	return passphraseHash, nil
}

func generatePassphrase(
	encrypt bcrypt.Encryptor,
	pass []byte,
) ([]byte, error) {

	if len(pass) == 0 {
		fmt.Println("Continue without passphrase")
		return nil, nil
	}

	passphraseHash, err := encrypt.Generate(pass)
	if err != nil {
		return pass, fmt.Errorf("could not encrypt passphrase: %v", err)
	}

	return passphraseHash, nil

}
