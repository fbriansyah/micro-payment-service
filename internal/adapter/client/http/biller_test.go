package httpclient

import (
	"testing"

	"github.com/fbriansyah/micro-payment-service/util"
	"github.com/stretchr/testify/require"
)

func TestInquiry(t *testing.T) {
	billNumber := "6310233333334"
	expectedBaseAmount := int64(10000)
	expectedFineAmount := int64(2000)
	expectedTotalAmount := int64(12000)

	response, err := testAdapter.Inquiry("6310233333334")
	require.NoError(t, err)
	require.NotEmpty(t, response)

	require.Equal(t, false, response.Error)
	require.Equal(t, billNumber, response.Data.BillNumber)

	require.Equal(t, expectedBaseAmount, response.Data.BaseAmount)
	require.Equal(t, expectedFineAmount, response.Data.FineAmount)
	require.Equal(t, expectedTotalAmount, response.Data.TotalAmount)
}

func TestInquiryNotFound(t *testing.T) {
	response, err := testAdapter.Inquiry("6310233333330")
	require.Error(t, ErrorBillNotFound, err)
	require.Empty(t, response)
}

func TestPayment(t *testing.T) {
	billNumber := "6310233333334"
	inqResponse, err := testAdapter.Inquiry(billNumber)
	require.NoError(t, err)
	require.NotEmpty(t, inqResponse)

	payResponse, err := testAdapter.Payment(PaymentRequest{
		BillNumber:       billNumber,
		RefferenceNumber: util.RandomRefferenceNumber(),
		TotalAmount:      inqResponse.Data.TotalAmount,
	})
	require.NoError(t, err)
	require.NotEmpty(t, payResponse)
}
