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
	Inquiry(ctx context.Context, arg dmrequest.InquryRequestParams) (dmlog.RequestLog, error)
	Payment(ctx context.Context, arg dmrequest.PaymentRequestParams) (dmtransaction.Transaction, error)
	GetUserBalance(ctx context.Context, userID uuid.UUID) (int64, error)
	GetAllProducts(ctx context.Context) ([]dmproduct.Product, error)
}
