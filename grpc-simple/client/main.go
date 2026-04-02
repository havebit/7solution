package main

import (
	"context"
	"fmt"
	"log"
	"time"

	helloproto "github.com/pallat/helloproto/buf/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := helloproto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloproto.HelloRequest{Name: "Pallat"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("response: %v", r.GetMessage())
}
