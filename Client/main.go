package main

import (
	"log"

	proto "github.com/Amizhthanmd/go_grpc/Proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Can't connect : %v", err)
	}

	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	names := &proto.NamesList{
		Names: []string{"Amizhthan","Mugesh","Mani","Manoj"},
	} 

	CallSayHello(client)
	callSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
