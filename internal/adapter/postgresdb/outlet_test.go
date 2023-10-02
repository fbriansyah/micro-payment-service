package postgresdb

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomOutlet(t *testing.T) Outlet {
	arg := CreateOutletParams{
		User:    uuid.New(),
		Deposit: 1000000,
	}

	outlet, err := testQueries.CreateOutlet(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, outlet)

	require.Equal(t, arg.User, outlet.User)
	require.Equal(t, arg.Deposit, outlet.Deposit)
	require.Equal(t, true, outlet.IsActive)

	return outlet
}

func TestCreateOutlet(t *testing.T) {
	CreateRandomOutlet(t)
}

func TestGetOutletByID(t *testing.T) {
	outlet1 := CreateRandomOutlet(t)

	outlet2, err := testQueries.GetOutletByUserID(context.Background(), outlet1.User)

	require.NoError(t, err)
	require.NotEmpty(t, outlet2)

	require.Equal(t, outlet1.ID, outlet2.ID)
	require.Equal(t, outlet1.Deposit, outlet2.Deposit)
}

func TestUpdateDeposit(t *testing.T) {
	outlet1 := CreateRandomOutlet(t)

	outlet2, err := testQueries.UpdateOutletDeposit(context.Background(), UpdateOutletDepositParams{
		Deposit: outlet1.Deposit + 500000,
		ID:      outlet1.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, outlet2)

	require.Equal(t, outlet1.Deposit+500000, outlet2.Deposit)
}
