package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/pallat/training/7solution/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type paymentServer struct {
	proto.UnimplementedPaymentServiceServer
	nodeID string
}

func (s *paymentServer) ProcessPayment(ctx context.Context, req *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	log.Printf("Node %s: Processing payment for order %s\n", s.nodeID, req.OrderId)
	return &proto.PaymentResponse{
		Status: "SUCCESS",
		NodeId: s.nodeID,
	}, nil
}

type orderServer struct {
	proto.UnimplementedOrderServiceServer
	paymentClient proto.PaymentServiceClient
}

func (s *orderServer) CreateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {
	log.Printf("Order Service: Creating order for %s x %d\n", req.ItemName, req.Quantity)

	orderID := fmt.Sprintf("ORDER-%d", time.Now().Unix())
	paymentRes, err := s.paymentClient.ProcessPayment(ctx, &proto.PaymentRequest{
		OrderId: orderID,
		Amount:  float32(req.Quantity) * 10.0,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to process payment: %v", err)
	}

	return &proto.OrderResponse{
		OrderId:       orderID,
		Status:        "CREATED",
		PaymentStatus: fmt.Sprintf("%s (Processed by %s)", paymentRes.Status, paymentRes.NodeId),
	}, nil
}

func main() {
	serviceType := os.Getenv("SERVICE_TYPE")
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	if serviceType == "payment" {
		nodeID := os.Getenv("NODE_ID")
		if nodeID == "" {
			hostname, _ := os.Hostname()
			nodeID = hostname
			if nodeID == "" {
				nodeID = "default"
			}
		}
		proto.RegisterPaymentServiceServer(s, &paymentServer{nodeID: nodeID})
		log.Printf("Payment Service (Node: %s) listening on :%s\n", nodeID, port)
	} else if serviceType == "order" {
		paymentAddr := os.Getenv("PAYMENT_ADDR")
		if paymentAddr == "" {
			paymentAddr = "localhost:50051"
		}
		conn, err := grpc.NewClient(paymentAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		paymentClient := proto.NewPaymentServiceClient(conn)
		proto.RegisterOrderServiceServer(s, &orderServer{paymentClient: paymentClient})
		log.Printf("Order Service listening on :%s\n", port)
	} else {
		log.Fatalf("unknown service type: %s", serviceType)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
