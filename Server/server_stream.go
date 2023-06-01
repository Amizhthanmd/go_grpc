package main

import (
	"log"
	"time"

	proto "github.com/Amizhthanmd/go_grpc/Proto"
)

func (s *helloServer) SayHelloServerStreaming(req *proto.NamesList, stream proto.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names : %v", req.Names)
	for _, name := range req.Names {
		res := &proto.HelloResponse{
			Message: "Hai " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}


