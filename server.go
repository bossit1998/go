package main

import (
	"log"
	"net"

	"github.com/bossit98/go_tutorial/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server over port 9000: %v", err)
	}
}
