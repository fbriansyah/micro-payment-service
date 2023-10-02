package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	num := RandomInt(10, 99)

	require.Greater(t, num, int64(9))
	require.Less(t, num, int64(100))
}

func TestRandomString(t *testing.T) {
	str := RandomString(9)

	require.Equal(t, len(str), 9)
}

func TestRandomMoney(t *testing.T) {
	money := RandomMoney()

	require.Greater(t, money, int64(999))
	require.Less(t, money, int64(100001))
}

func TestRandomBillNumber(t *testing.T) {
	billNumb := RandomBillNumber()

	require.Equal(t, billNumb[:2], "63")
	require.Equal(t, len(billNumb), 18)
}

func TestRandomRefferenceNumber(t *testing.T) {
	reff := RandomRefferenceNumber()

	require.Equal(t, len(reff), 40)
}
