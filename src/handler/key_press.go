package handler

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/elbuki/ctrl-api/src/control"

	pb "github.com/elbuki/ctrl-protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) KeyPress(
	ctx context.Context,
	req *pb.KeyPressRequest,
) (*empty.Empty, error) {

	key, err := control.TranslateProtoKey(req.GetKey())
	if err != nil {
		gErr := status.Error(
			codes.FailedPrecondition,
			fmt.Sprintf("could not translate key: %v", err),
		)
		return nil, gErr
	}

	if err := h.api.conf.Controller.SendKeys(key); err != nil {
		gErr := status.Error(
			codes.Unknown,
			fmt.Sprintf("could not send key from external: %v", err),
		)
		return nil, gErr
	}

	return new(empty.Empty), nil
}
