package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/bossit1998/go/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "hello from client",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when calling say hello: %s", err)
	}
	log.Printf("response from server: %s", response.Body)
}
