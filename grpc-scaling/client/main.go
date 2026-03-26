package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/pallat/training/7solution/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Point to Nginx LB (port 80)
	addr := os.Getenv("ORDER_ADDR")
	if addr == "" {
		addr = "localhost:80"
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := c.CreateOrder(ctx, &proto.OrderRequest{
		ItemName: "Scalable Gadget",
		Quantity: 5,
	})
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}

	log.Printf("Order Response: ID=%s, Status=%s, Payment=%s", r.OrderId, r.Status, r.PaymentStatus)
}
