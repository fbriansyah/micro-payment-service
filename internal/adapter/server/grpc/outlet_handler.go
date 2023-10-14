package grpcserver

import (
	"context"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/payment"
	"github.com/google/uuid"
)

// GetBalance serve as grpc server handler, if any error it wills return -1 balance
func (a *GrpcServerAdapter) GetBalance(ctx context.Context, req *payment.GetBalanceRequest) (*payment.GetBalanceResponse, error) {
	balance, err := a.paymentService.GetUserBalance(ctx, uuid.MustParse(req.UserId))
	if err != nil {
		return &payment.GetBalanceResponse{Balance: -1}, err
	}

	return &payment.GetBalanceResponse{
		Balance: float64(balance),
	}, nil
}
