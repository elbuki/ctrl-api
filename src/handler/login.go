package handler

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type LoginRequest struct {
	UUID       string
	Passphrase []byte
}

type LoginResponse struct {
	Token []byte
}

func (a *API) Pair(req *LoginRequest, res *LoginResponse) error {
	var err error

	if req.UUID == "" {
		return fmt.Errorf("could not receive the uuid from device")
	}

	if a.Conf.UsePassphrase {
		err = a.Conf.Encryptor.Compare(a.SavedPassphrase, req.Passphrase)
	}

	if err != nil {
		return fmt.Errorf("could not see the match for passphrases: %v", err)
	}

	res.Token, err = generateToken(req.UUID)
	if err != nil {
		return fmt.Errorf("could not generate token: %v", err)
	}

	return nil
}

func generateToken(uuid string) ([]byte, error) {
	h := sha256.New()
	iso8601Date := time.Now().Format(time.RFC3339)
	plainToken := uuid + "_" + iso8601Date

	if _, err := h.Write([]byte(plainToken)); err != nil {
		return nil, fmt.Errorf("could not write to hash: %v", err)
	}

	return h.Sum(nil), nil
}
