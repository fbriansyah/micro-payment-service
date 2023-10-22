package postgresdb

import (
	"context"
	"testing"

	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/stretchr/testify/require"
)

func TestPayment(t *testing.T) {

	product := CreateRandomProduct(t)
	outlet := CreateRandomOutlet(t)

	logInq := CreateRandomRequestLog(t, "INQ", util.RandomMoney(), product, outlet)

	arg := PaymentParam{
		LogInquiry:       logInq,
		UserId:           outlet.User,
		RefferenceNumber: util.RandomRefferenceNumber(),
		PayResponseStr:   "{}",
	}

	tx, logPay, err := testAdapter.PaymentTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tx)
	require.NotEmpty(t, logPay)

	outAfterPay, err := testQueries.GetOutletByUserID(context.Background(), outlet.User)
	require.NoError(t, err)

	require.Equal(t, outlet.Deposit-logInq.TotalAmount, outAfterPay.Deposit)
}
