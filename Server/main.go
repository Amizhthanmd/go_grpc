package main

import (
	"log"
	"net"

	proto "github.com/Amizhthanmd/go_grpc/Proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Can't start a server : %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server Starting at : %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Can't start : %v", err)
	}
}