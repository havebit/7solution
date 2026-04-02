package main

import (
	"context"
	"log"
	"net"
	"os"

	helloproto "github.com/pallat/helloproto/buf/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	helloproto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *helloproto.HelloRequest) (*helloproto.HelloResponse, error) {
	log.Printf("get: %v\n", in.GetName())

	return &helloproto.HelloResponse{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	helloproto.RegisterGreeterServer(s, &server{})

	if os.Getenv("APP_ENV") != "production" {
		reflection.Register(s)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
