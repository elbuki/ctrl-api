package handler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	pb "github.com/elbuki/ctrl-protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Login(
	ctx context.Context,
	req *pb.LoginRequest,
) (*pb.LoginResponse, error) {

	var err error

	if req.GetUuid() == "" {
		gErr := status.Error(
			codes.InvalidArgument,
			"could not receive the uuid from device",
		)
		return nil, gErr
	}

	if h.api.conf.UsePassphrase {
		err = h.api.conf.Encryptor.Compare(
			h.api.savedPassphrase,
			req.GetPassphrase(),
		)
	}

	if err != nil {
		gErr := status.Error(
			codes.FailedPrecondition,
			fmt.Sprintf("could not match passphrases: %v", err),
		)
		return nil, gErr
	}

	token, err := generateToken(req.GetUuid())
	if err != nil {
		gErr := status.Error(
			codes.Unknown,
			fmt.Sprintf("could not generate token: %v", err),
		)
		return nil, gErr
	}

	h.api.token = token

	return &pb.LoginResponse{Token: token}, nil

}

func generateToken(uuid string) (string, error) {
	h := sha256.New()
	iso8601Date := time.Now().Format(time.RFC3339)
	plainToken := uuid + "_" + iso8601Date

	if _, err := h.Write([]byte(plainToken)); err != nil {
		return "", fmt.Errorf("could not write to hash: %v", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
