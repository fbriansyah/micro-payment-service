package port

import (
	"context"

	dmlog "github.com/fbriansyah/micro-payment-service/internal/application/domain/log"
	dmrequest "github.com/fbriansyah/micro-payment-service/internal/application/domain/request"
)

type PaymetService interface {
	Inquiry(ctx context.Context, arg dmrequest.InquryRequestParams) (dmlog.RequestLog, error)
	Payment(ctx context.Context, arg dmrequest.PaymentRequestParams)
}
