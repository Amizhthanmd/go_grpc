package main

import (
	"log"

	proto "github.com/Amizhthanmd/go_grpc/Proto"
)

func (s *helloServer) SayHelloClientStreaming(stream proto.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err != nil {
			return stream.SendAndClose(&proto.MessagesList{Messages: messages})
		}
		if err != nil {
			return nil
		}
		log.Printf("Got the request %v", req.Name)
		messages = append(messages, "Hello " + req.Name)
	}
}


