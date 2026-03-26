package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/pallat/training/chapters/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := os.Getenv("ORDER_ADDR")
	if addr == "" {
		addr = "localhost:50051"
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ProcessPayment(ctx, &proto.PaymentRequest{
		OrderId: "01",
	})
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}

	log.Printf("Order Response: ID=%s, Status=%s, Payment=%s", r.OrderId, r.Status, r.PaymentStatus)
}
