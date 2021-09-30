package handler

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func WithInterceptor(h *Handler) grpc.ServerOption {
	return grpc.UnaryInterceptor(AuthorizationInterceptor(h))
}

func AuthorizationInterceptor(h *Handler) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		ignoredMethods := []string{"/pb.MainService/Login"}
		for _, m := range ignoredMethods {
			if m == info.FullMethod {
				return handler(ctx, req)
			}
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			err := status.Error(
				codes.FailedPrecondition,
				"could not parse metadata",
			)
			return nil, err
		}

		authMD := md.Get("Authorization")
		if len(authMD) == 0 {
			err := status.Error(
				codes.FailedPrecondition,
				"could not get auth token",
			)
			return nil, err
		}

		if len(authMD) == 0 {
			err := status.Error(
				codes.FailedPrecondition,
				"could not find a token in the metadata",
			)
			return nil, err
		}

		if h.api.token == "" {
			err := status.Error(
				codes.FailedPrecondition,
				"could not get passed token, not logged in",
			)
			return nil, err
		}

		if h.api.token != authMD[0] {
			err := status.Error(
				codes.FailedPrecondition,
				"could not verify token",
			)
			return nil, err
		}

		return handler(ctx, req)

	}
}
