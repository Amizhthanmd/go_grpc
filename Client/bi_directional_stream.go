package main

import (
	"context"
	"io"
	"log"
	"time"

	proto "github.com/Amizhthanmd/go_grpc/Proto"
)

func callSayHelloBidirectionalStream(client proto.GreetServiceClient, names *proto.NamesList) {

	log.Printf("Bidirectional Streaming started...")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	ch := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("bistream error : %v", err)
			}
			log.Println(message)
		}
		close(ch)
	}()

	for _, name := range names.Names {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-ch
	log.Printf("Bidirectional Streaming finished")
}
