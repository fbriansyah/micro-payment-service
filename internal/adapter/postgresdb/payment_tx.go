package postgresdb

import (
	"context"

	dmlog "github.com/fbriansyah/micro-payment-service/internal/application/domain/log"
	dmtransaction "github.com/fbriansyah/micro-payment-service/internal/application/domain/transaction"
	"github.com/google/uuid"
)

type PaymentParam struct {
	LogInquiry       RequestLog
	UserId           uuid.UUID
	RefferenceNumber string
	PayResponseStr   string
}

func (a *DatabaseStore) PaymentTx(ctx context.Context, arg PaymentParam) (dmtransaction.Transaction, error) {
	var transaction dmtransaction.Transaction

	err := a.execTx(ctx, func(q *Queries) error {
		logPay, err := q.CreateRequestLog(ctx, CreateRequestLogParams{
			Mode:           dmlog.SERVICE_MODE_PAY,
			Product:        arg.LogInquiry.Product,
			BillNumber:     arg.LogInquiry.BillNumber,
			Name:           arg.LogInquiry.Name,
			TotalAmount:    arg.LogInquiry.TotalAmount,
			Outlet:         arg.LogInquiry.Outlet,
			BillerResponse: arg.PayResponseStr,
		})

		if err != nil {
			return err
		}
		trx, err := q.CreateTransaction(ctx, CreateTransactionParams{
			BillNumber:       logPay.BillNumber,
			Product:          logPay.Product,
			Inquiry:          uuid.NullUUID{Valid: true, UUID: arg.LogInquiry.ID},
			Payment:          uuid.NullUUID{Valid: true, UUID: logPay.ID},
			Amount:           logPay.TotalAmount,
			RefferenceNumber: arg.RefferenceNumber,
			Outlet:           arg.LogInquiry.Outlet,
			Status:           1,
		})
		if err != nil {
			return err
		}

		outlet, err := q.GetOutletByUserID(ctx, arg.UserId)
		if err != nil {
			return err
		}

		_, err = q.UpdateOutletDeposit(ctx, UpdateOutletDepositParams{
			Deposit: outlet.Deposit - trx.Amount,
			ID:      outlet.ID,
		})
		if err != nil {
			return err
		}

		transaction.Amount = trx.Amount
		transaction.BillNumber = trx.BillNumber
		transaction.ID = trx.ID
		transaction.Name = logPay.Name
		transaction.RefferenceNumber = arg.RefferenceNumber
		transaction.TransactionDatetime = trx.TransactionDatetime

		return nil
	})

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
