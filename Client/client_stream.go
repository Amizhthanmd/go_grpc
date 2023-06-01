package main

import (
	"context"
	"log"
	"time"

	proto "github.com/Amizhthanmd/go_grpc/Proto"
)

func callSayHelloClientStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Client Streaming Started...")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could'nt send names %v", err)
	}
	for _, name := range names.Names {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished...")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
