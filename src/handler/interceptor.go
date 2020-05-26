package handler

import (
	"bytes"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
			return nil, fmt.Errorf("could not parse metadata")
		}

		authMD := md.Get("Authorization")
		if len(authMD) == 0 {
			return nil, fmt.Errorf("could not get auth token")
		}

		passedToken := []byte(authMD[0])

		if len(h.api.token) == 0 {
			return nil, fmt.Errorf("could not get passed token, not logged in")
		}

		if !bytes.Equal(h.api.token, passedToken) {
			return nil, fmt.Errorf("could not verify token")
		}

		return handler(ctx, req)

	}
}
