package handler

import (
	"context"
	"fmt"

	"github.com/elbuki/ctrl-api/src/control"

	pb "github.com/elbuki/ctrl-protobuf/src/golang"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) KeyPress(
	ctx context.Context,
	req *pb.KeyPressRequest,
) (*empty.Empty, error) {

	key, err := control.TranslateProtoKey(req.GetKey())
	if err != nil {
		return nil, fmt.Errorf("could not translate key: %v", err)
	}

	if err := h.api.conf.Controller.SendKeys(key); err != nil {
		return nil, fmt.Errorf("could not send key from external: %v", err)
	}

	return nil, nil
}
