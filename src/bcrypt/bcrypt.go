package bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Encryptor struct {
	hashCost int
}

func NewEncryptor(cost int) Encryptor {
	return Encryptor{hashCost: cost}
}

func (e Encryptor) Compare(hashed, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashed, password)
	if err != nil {
		return fmt.Errorf("could not compare hash with password: %v", err)
	}

	return nil
}

func (e Encryptor) Generate(password []byte) ([]byte, error) {
	pass, err := bcrypt.GenerateFromPassword(password, e.hashCost)
	if err != nil {
		return nil, fmt.Errorf("could not generate password: %v", err)
	}

	return pass, nil
}
