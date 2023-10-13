package grpcserver

import (
	"context"
	"fmt"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/payment"
	dmrequest "github.com/fbriansyah/micro-payment-service/internal/application/domain/request"
	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

// Inquiry implementation of grpc server
func (a *GrpcServerAdapter) Inquiry(ctx context.Context, reqParam *payment.InquiryRequest) (*payment.InquiryResponse, error) {
	logInq, err := a.paymentService.Inquiry(ctx, dmrequest.InquryRequestParams{
		BillNumber: reqParam.BillNumber,
		UserId:     uuid.MustParse(reqParam.UserId),
		Product:    reqParam.ProductCode,
	})
	if err != nil {
		return &payment.InquiryResponse{}, generateError(
			codes.FailedPrecondition,
			fmt.Sprintf("error inquiry: %v", err),
		)
	}

	return &payment.InquiryResponse{
		InqId:       logInq.ID.String(),
		BillNumber:  logInq.BillNumber,
		ProductCode: logInq.Product,
		Name:        logInq.Name,
		TotalAmount: float64(logInq.TotalAmount),
		DetailBill: &payment.DetailBill{
			BaseAmount: float64(logInq.BaseAmount),
			FineAmount: float64(logInq.FineAmount),
		},
	}, nil
}

// Payment implementation of grpc server
func (a *GrpcServerAdapter) Payment(ctx context.Context, reqParam *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	trx, err := a.paymentService.Payment(ctx, dmrequest.PaymentRequestParams{
		UserId: uuid.MustParse(reqParam.UserId),
		LogInq: uuid.MustParse(reqParam.InqId),
	})
	if err != nil {
		return &payment.PaymentResponse{}, generateError(
			codes.FailedPrecondition,
			fmt.Sprintf("error payment: %v", err),
		)
	}

	return &payment.PaymentResponse{
		BillNumber:          trx.BillNumber,
		ProductCode:         trx.Product,
		Name:                trx.Name,
		TotalAmount:         float64(trx.Amount),
		RefferenceNumber:    trx.RefferenceNumber,
		TransactionDatetime: util.ToDateTime(trx.TransactionDatetime),
	}, nil
}
