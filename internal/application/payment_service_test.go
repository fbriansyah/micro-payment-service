package application

import (
	"context"
	"fmt"
	"testing"

	"github.com/fbriansyah/micro-payment-service/internal/adapter/postgresdb"
	dmrequest "github.com/fbriansyah/micro-payment-service/internal/application/domain/request"
	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestInquiry(t *testing.T) {
	product, err := testQueries.CreateProduct(
		context.Background(),
		postgresdb.CreateProductParams{
			ProductCode: util.RandomString(5),
			ProductName: fmt.Sprintf(
				"%s %s",
				util.RandomString(4),
				util.RandomString(5),
			),
		},
	)
	require.NoError(t, err)
	outlet, err := testQueries.CreateOutlet(
		context.Background(),
		postgresdb.CreateOutletParams{
			User:    uuid.New(),
			Deposit: 999999,
		},
	)
	require.NoError(t, err)

	arg := dmrequest.InquryRequestParams{
		BillNumber: "6310233333334",
		UserId:     outlet.User,
		Product:    product.ProductCode,
	}
	inqLog, err := testService.Inquiry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, inqLog)

	require.Equal(t, arg.BillNumber, inqLog.BillNumber)
	require.Equal(t, arg.Product, inqLog.Product)
}

func TestPayment(t *testing.T) {
	product, err := testQueries.CreateProduct(
		context.Background(),
		postgresdb.CreateProductParams{
			ProductCode: util.RandomString(5),
			ProductName: fmt.Sprintf(
				"%s %s",
				util.RandomString(4),
				util.RandomString(5),
			),
		},
	)
	require.NoError(t, err)
	outletParam := postgresdb.CreateOutletParams{
		User:    uuid.New(),
		Deposit: 999999,
	}
	outlet, err := testQueries.CreateOutlet(
		context.Background(),
		outletParam,
	)
	require.NoError(t, err)

	arg := dmrequest.InquryRequestParams{
		BillNumber: "6310233333335",
		UserId:     outlet.User,
		Product:    product.ProductCode,
	}
	inqLog, err := testService.Inquiry(context.Background(), arg)
	require.NoError(t, err)

	argPay := dmrequest.PaymentRequestParams{
		UserId: outlet.User,
		LogInq: inqLog.ID,
	}

	trx, err := testService.Payment(
		context.Background(),
		argPay,
	)
	require.NoError(t, err)
	require.NotEmpty(t, trx)

	require.NotEqual(t, "", trx.RefferenceNumber)

	outletAfterUpdate, err := testQueries.GetOutletByUserID(context.Background(), arg.UserId)
	require.NoError(t, err)
	require.Equal(t, outletParam.Deposit-trx.Amount, outletAfterUpdate.Deposit)
}
