package handler

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/elbuki/ctrl-api/src/pb"
)

type LoginHandler struct {
	api *API
}

func NewLoginHandler(a *API) *LoginHandler {
	return &LoginHandler{a}
}

func (h *LoginHandler) Login(
	ctx context.Context,
	req *pb.LoginRequest,
) (*pb.LoginResponse, error) {

	var token []byte
	var err error

	if req.GetUuid() == "" {
		return nil, fmt.Errorf("could not receive the uuid from device")
	}

	if h.api.conf.UsePassphrase {
		err = h.api.conf.Encryptor.Compare(
			h.api.savedPassphrase,
			req.GetPassphrase(),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("could not match passphrases: %v", err)
	}

	if h.api.conf.UsePassphrase {
		token, err = generateToken(req.GetUuid())
	}

	if err != nil {
		return nil, fmt.Errorf("could not generate token: %v", err)
	}

	return &pb.LoginResponse{Token: token}, nil

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
