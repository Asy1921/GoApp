package sqlc

import (
	"backend/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestCreateAccount(t *testing.T) {
	arg:= CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),

	}
	account, err:= testQueries.CreateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,account)
	// require.Equal(t. .Owner,account.Owner)
	require.NotZero(t,account.ID)
	require.NotZero(t,account.CreatedAt)
}

func TestGetOwnerByID(t *testing.T) {

	account, err:= testQueries.GetAccount(context.Background(),4)
	t.Logf(account.Owner)
	t.Log(err)
}

