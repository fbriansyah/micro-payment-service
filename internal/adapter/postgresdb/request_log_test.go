package postgresdb

import (
	"context"
	"testing"

	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRequestLog(
	t *testing.T, mode string, product Product, outlet Outlet) RequestLog {
	arg := CreateRequestLogParams{
		Mode:           mode,
		Product:        product.ProductCode,
		BillNumber:     util.RandomBillNumber(),
		Name:           util.RandomString(5),
		TotalAmount:    util.RandomMoney(),
		BillerResponse: util.RandomString(10),
		Outlet:         outlet.ID,
	}

	log, err := testQueries.CreateRequestLog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, log)

	require.Equal(t, arg.Mode, log.Mode)
	require.Equal(t, arg.Product, log.Product)
	require.Equal(t, arg.BillNumber, log.BillNumber)
	require.Equal(t, arg.Name, log.Name)
	require.Equal(t, arg.TotalAmount, log.TotalAmount)
	require.Equal(t, arg.BillerResponse, log.BillerResponse)

	return log
}

func TestCreateRequestLog(t *testing.T) {
	product := CreateRandomProduct(t)
	outlet := CreateRandomOutlet(t)

	CreateRandomRequestLog(t, "INQ", product, outlet)
}

func TestGetRequestLogByID(t *testing.T) {
	product := CreateRandomProduct(t)
	outlet := CreateRandomOutlet(t)

	log1 := CreateRandomRequestLog(t, "INQ", product, outlet)

	log2, err := testQueries.GetRequestLogByID(context.Background(), log1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, log2)

	require.Equal(t, "INQ", log1.Mode)
	require.Equal(t, log1.BillNumber, log2.BillNumber)
}
