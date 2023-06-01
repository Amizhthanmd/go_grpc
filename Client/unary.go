package main

import (
	"context"
	"log"
	"time"

	proto "github.com/Amizhthanmd/go_grpc/Proto"
)

func CallSayHello(client proto.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatalf("can't send : %v", err)
	}
	log.Printf("%s", res.Message)
}
