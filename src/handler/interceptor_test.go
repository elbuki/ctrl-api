package handler_test

import (
	"context"
	"testing"

	"github.com/elbuki/ctrl-api/src/config"

	"github.com/elbuki/ctrl-api/src/handler"
	pb "github.com/elbuki/ctrl-protobuf/proto"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type authInterceptorScenario struct {
	ctx    context.Context
	info   *grpc.UnaryServerInfo
	h      *handler.Handler
	throws bool
}

func TestAuthorizationInterceptor(t *testing.T) {
	ctx := context.Background()
	conf := config.Config{}
	info := &grpc.UnaryServerInfo{
		FullMethod: "TestService.AuthorizationInterceptor",
	}
	loginInfo := &grpc.UnaryServerInfo{FullMethod: "/pb.MainService/Login"}
	invalidMetadata := metadata.NewIncomingContext(
		ctx, metadata.Pairs("foo", "bar"),
	)
	validMetadata := metadata.NewIncomingContext(
		ctx, metadata.Pairs("Authorization", "test_passphrase"),
	)
	tH := handler.NewHandler(conf, nil)
	differentToken := handler.NewHandler(conf, nil)
	happyPathToken := handler.NewHandler(conf, nil)
	unaryHandler := func(
		ctx context.Context,
		req interface{},
	) (interface{}, error) {
		return nil, nil
	}

	_, err := differentToken.Login(
		ctx, &pb.LoginRequest{Uuid: "foo", Passphrase: []byte("bar")},
	)

	if err != nil {
		t.Errorf("could not login for getting a different token: %v", err)
	}

	res, err := happyPathToken.Login(
		ctx,
		&pb.LoginRequest{Uuid: "test", Passphrase: []byte("test_passphrase")},
	)

	if err != nil {
		t.Errorf("could not login for happy path: %v", err)
	}

	happyPathMetadata := metadata.NewIncomingContext(
		ctx, metadata.Pairs("Authorization", string(res.GetToken())),
	)

	var table = []authInterceptorScenario{
		{ctx: ctx, info: info, h: tH, throws: true},
		{ctx: invalidMetadata, info: info, h: tH, throws: true},
		{ctx: validMetadata, info: info, h: tH, throws: true},
		{ctx: validMetadata, info: info, h: differentToken, throws: true},
		{ctx: validMetadata, info: loginInfo, h: happyPathToken},
		{ctx: happyPathMetadata, info: info, h: happyPathToken},
	}

	for i, s := range table {
		ai := handler.AuthorizationInterceptor(s.h)
		_, err := ai(s.ctx, &pb.KeyPressRequest{}, s.info, unaryHandler)
		if !s.throws && err != nil {
			t.Errorf("%d: could not execute auth interceptor: %v", i, err)
		}

		if s.throws && err == nil {
			t.Errorf("%d: expected error but it went fine", i)
		}
	}
}
