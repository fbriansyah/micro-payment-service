package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/payment"
	"github.com/fbriansyah/micro-payment-service/internal/port"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServerAdapter struct {
	paymentService port.PaymentService
	server         *grpc.Server
	grpcPort       int

	payment.UnimplementedPaymentServiceServer
}

func NewGrpcServerAdapter(paymentService port.PaymentService, grpcPort int) *GrpcServerAdapter {
	return &GrpcServerAdapter{
		paymentService: paymentService,
		grpcPort:       grpcPort,
	}
}

func (a *GrpcServerAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("failed to listen on port %d: %v\n", a.grpcPort, err)
	}
	log.Printf("Server listen on port %d \n", a.grpcPort)

	grpcServer := grpc.NewServer()

	a.server = grpcServer

	payment.RegisterPaymentServiceServer(grpcServer, a)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server grpc on port %d: %v\n", a.grpcPort, err)
	}
}

// Stop the grpc server
func (a *GrpcServerAdapter) Stop() {
	a.server.Stop()
}

func generateError(code codes.Code, msg string) error {
	s := status.New(code, msg)
	s, _ = s.WithDetails(&errdetails.ErrorInfo{})
	return s.Err()
}
