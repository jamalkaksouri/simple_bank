package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jamalkaskouri/simple_bank/internal/db/util"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	generatedAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, generatedAccount)
	require.Equal(t, newAccount.ID, generatedAccount.ID)
	require.Equal(t, newAccount.Owner, generatedAccount.Owner)
	require.Equal(t, newAccount.Balance, generatedAccount.Balance)
	require.Equal(t, newAccount.Currency, generatedAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, generatedAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: util.RandomMoney(),
	}

	generatedAccount, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, generatedAccount)
	require.Equal(t, newAccount.ID, generatedAccount.ID)
	require.Equal(t, newAccount.Owner, generatedAccount.Owner)
	require.Equal(t, args.Balance, generatedAccount.Balance)
	require.Equal(t, newAccount.Currency, generatedAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, generatedAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	generatedAccount, errGet := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.Error(t, errGet)
	require.EqualError(t, errGet, sql.ErrNoRows.Error())
	require.Empty(t, generatedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
