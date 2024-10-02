package sqlc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestCreateAccount(t *testing.T) {
	arg:= CreateAccountParams{
		Owner: "Amitabh",
		Balance: 1000,
		Currency: "INR",

	}
	account, err:= testQueries.CreateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,account)
	// require.Equal(t. .Owner,account.Owner)
	require.NotZero(t,account.ID)
	require.NotZero(t,account.CreatedAt)
}

