package postgresdb

import (
	"context"
	"testing"

	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransaction(t *testing.T, product Product, outlet Outlet) Transaction {
	logInq := CreateRandomRequestLog(
		t,
		"INQ",
		util.RandomMoney(),
		product,
		outlet,
	)
	logPay := CreateRandomRequestLog(
		t,
		"PAY",
		logInq.TotalAmount,
		product,
		outlet,
	)

	arg := CreateTransactionParams{
		BillNumber: util.RandomBillNumber(),
		Product:    product.ProductCode,
		Outlet:     outlet.ID,
		Inquiry: uuid.NullUUID{
			Valid: true,
			UUID:  logInq.ID,
		},
		Payment: uuid.NullUUID{
			Valid: true,
			UUID:  logPay.ID,
		},
		Amount:           logInq.TotalAmount,
		RefferenceNumber: util.RandomRefferenceNumber(),
		Status:           1,
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, transaction.Amount, logInq.TotalAmount)
	require.Equal(t, arg.RefferenceNumber, transaction.RefferenceNumber)
	require.Equal(t, logInq.ID, transaction.Inquiry.UUID)
	require.Equal(t, logPay.ID, transaction.Payment.UUID)
	require.Equal(t, arg.Product, transaction.Product)
	require.Equal(t, arg.Outlet, transaction.Outlet)
	require.Equal(t, int32(1), transaction.Status)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	product := CreateRandomProduct(t)
	outlet := CreateRandomOutlet(t)

	CreateRandomTransaction(t, product, outlet)
}

func TestGetTransactionByID(t *testing.T) {
	product := CreateRandomProduct(t)
	outlet := CreateRandomOutlet(t)

	trx1 := CreateRandomTransaction(t, product, outlet)

	trx2, err := testQueries.GetTransactionByID(context.Background(), trx1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trx2)

	require.Equal(t, trx1.ID, trx2.ID)
	require.Equal(t, trx1.RefferenceNumber, trx2.RefferenceNumber)
	require.Equal(t, trx1.Amount, trx2.Amount)
}
