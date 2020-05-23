package handler

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func WithAuthorization() grpc.ServerOption {
	return grpc.UnaryInterceptor(auth)
}

func auth(
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

	// TODO: Validate token from the metadata

	log.Printf("%#v", info)

	return handler(ctx, req)

}
