package application

import (
	dmlog "github.com/fbriansyah/micro-payment-service/internal/application/domain/log"
	dmtransaction "github.com/fbriansyah/micro-payment-service/internal/application/domain/transaction"
	"github.com/google/uuid"
)

type PaymentService struct{}

type InquryRequestParams struct {
	BillNumber string
	UserId     uuid.UUID
	Product    string
}

func (s *PaymentService) Inquiry(arg InquryRequestParams) (dmlog.RequestLog, error) {
	// 1. Get product endpoint from database.
	// 1. Send Inquiry Request to product biller service.
	// 1. Check response error
	// 		- if no error, save response to `request_logs`.
	// 		- if error, send error to broker
	// 1. Send inqury response to broker (`bill data`, `request_logs.id`).

	return dmlog.RequestLog{}, nil
}

type PaymentRequestParams struct {
	UserId      uuid.UUID
	BillNumber  string
	Amount      int64
	ProductCode string
}

func (s *PaymentService) Payment(arg PaymentRequestParams) (dmtransaction.Transaction, error) {
	// 1. Get Payment request from broker, data:
	// 		- `kiosks.id`
	// 		- inquiry `request_logs.id`
	// 		- `bill_number`
	// 		- `amount`
	// 		- `product_code`
	// 1. Check `kiosks.deposit`:
	// 		- if `deposit` > `amount` continue.
	// 		- if `deposit` < `amount` response error.
	// 1. Get product endpoint from database.
	// 1. Send payment request to biller.
	// 1. Check response error.
	// 		- if no error, save response to `request_logs`.
	// 		- if error, send error to broker.
	// 1. Insert data payment to `transactions` table.
	// 1. Decrease `kiosks.deposit`.
	// 1. Response payment success to broker.

	return dmtransaction.Transaction{}, nil
}
