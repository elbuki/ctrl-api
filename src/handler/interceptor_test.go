package handler_test

import (
	"context"
	"testing"

	"github.com/elbuki/ctrl-api/src/config"

	"github.com/elbuki/ctrl-api/src/handler"
	pb "github.com/elbuki/ctrl-protobuf/src/golang"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type authInterceptorScenario struct {
	ctx         context.Context
	info        *grpc.UnaryServerInfo
	h           *handler.Handler
	shouldThrow bool
}

func TestAuthorizationInterceptor(t *testing.T) {
	// Invalid context
	// Invalid metadata
	// Empty saved token

	// Different token
	// Login endpoint
	// Happy path
	conf := config.Config{}
	info := &grpc.UnaryServerInfo{
		FullMethod: "TestService.AuthorizationInterceptor",
	}
	invalidMetadata := metadata.NewIncomingContext(
		context.Background(),
		metadata.Pairs("foo", "bar"),
	)
	validMetadata := metadata.NewIncomingContext(
		context.Background(),
		metadata.Pairs("Authorization", "test_passphrase"),
	)
	unaryHandler := func(
		ctx context.Context,
		req interface{},
	) (interface{}, error) {

		t.Log("handler executed")
		return nil, nil

	}

	var table = []authInterceptorScenario{
		authInterceptorScenario{
			ctx:         context.Background(),
			info:        info,
			h:           handler.NewHandler(conf),
			shouldThrow: true,
		},
		authInterceptorScenario{
			ctx:         invalidMetadata,
			info:        info,
			h:           handler.NewHandler(conf),
			shouldThrow: true,
		},
		authInterceptorScenario{
			ctx:         validMetadata,
			info:        info,
			h:           handler.NewHandler(conf),
			shouldThrow: true,
		},
		// authInterceptorScenario{
		// 	ctx:         validMetadata,
		// 	info:        info,
		// 	h:           handler.NewHandler(),
		// 	shouldThrow: true,
		// },
	}

	for _, s := range table {
		i := handler.AuthorizationInterceptor(s.h)
		_, err := i(s.ctx, &pb.KeyPressRequest{}, s.info, unaryHandler)
		if !s.shouldThrow && err != nil {
			t.Errorf("could not execute auth interceptor: %v", err)
		}

		if s.shouldThrow && err == nil {
			t.Error("expected error but it went fine")
		}
	}

}
