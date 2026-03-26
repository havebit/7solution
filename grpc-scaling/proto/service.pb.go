// Hand-written simplified proto implementation for demo purposes
package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
)

type OrderRequest struct {
	ItemName string `json:"item_name,omitempty"`
	Quantity int32  `json:"quantity,omitempty"`
}

func (*OrderRequest) Reset()         {}
func (*OrderRequest) String() string { return "" }
func (*OrderRequest) ProtoMessage()  {}

type OrderResponse struct {
	OrderId       string `json:"order_id,omitempty"`
	Status        string `json:"status,omitempty"`
	PaymentStatus string `json:"payment_status,omitempty"`
}

func (*OrderResponse) Reset()         {}
func (*OrderResponse) String() string { return "" }
func (*OrderResponse) ProtoMessage()  {}

type PaymentRequest struct {
	OrderId string  `json:"order_id,omitempty"`
	Amount  float32 `json:"amount,omitempty"`
}

func (*PaymentRequest) Reset()         {}
func (*PaymentRequest) String() string { return "" }
func (*PaymentRequest) ProtoMessage()  {}

type PaymentResponse struct {
	Status string `json:"status,omitempty"`
	NodeId string `json:"node_id,omitempty"`
}

func (*PaymentResponse) Reset()         {}
func (*PaymentResponse) String() string { return "" }
func (*PaymentResponse) ProtoMessage()  {}

// Service interfaces and clients

type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/proto.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type OrderServiceServer interface {
	CreateOrder(context.Context, *OrderRequest) (*OrderResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, nil
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OrderRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(OrderServiceServer).CreateOrder(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/proto.OrderService/CreateOrder",
				}
				handler := func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(OrderServiceServer).CreateOrder(ctx, req.(*OrderRequest))
				}
				return interceptor(ctx, in, info, handler)
			},
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}

type PaymentServiceClient interface {
	ProcessPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) ProcessPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/proto.PaymentService/ProcessPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type PaymentServiceServer interface {
	ProcessPayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) ProcessPayment(context.Context, *PaymentRequest) (*PaymentResponse, error) {
	return nil, nil
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPayment",
			Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PaymentRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil {
					return srv.(PaymentServiceServer).ProcessPayment(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/proto.PaymentService/ProcessPayment",
				}
				handler := func(ctx context.Context, req interface{}) (interface{}, error) {
					return srv.(PaymentServiceServer).ProcessPayment(ctx, req.(*PaymentRequest))
				}
				return interceptor(ctx, in, info, handler)
			},
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
