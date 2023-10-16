package port

import (
	"context"

	dmlog "github.com/fbriansyah/micro-payment-service/internal/application/domain/log"
	dmproduct "github.com/fbriansyah/micro-payment-service/internal/application/domain/product"
	dmrequest "github.com/fbriansyah/micro-payment-service/internal/application/domain/request"
	dmtransaction "github.com/fbriansyah/micro-payment-service/internal/application/domain/transaction"
	"github.com/google/uuid"
)

type PaymentService interface {
	// Inquiry call inquiry method in client adapter
	Inquiry(ctx context.Context, arg dmrequest.InquryRequestParams) (dmlog.RequestLog, error)
	// Payment call payment method in client and create payment log
	Payment(ctx context.Context, arg dmrequest.PaymentRequestParams) (dmtransaction.Transaction, error)
	// GetUserBalance get current user balance or deposit
	GetUserBalance(ctx context.Context, userID uuid.UUID) (int64, error)
	// Get all active product
	GetAllProducts(ctx context.Context) ([]dmproduct.Product, error)
}
