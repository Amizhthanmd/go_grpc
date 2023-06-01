package main

import (
	"context"

	proto "github.com/Amizhthanmd/go_grpc/Proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello, Welcome to go_grpc",
	}, nil
}

