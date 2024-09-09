package db

import (
	"context"
	"testing"

	"github.com/prakhar30/bankoindiana/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount() (CreateAccountParams, Account, error) {
	hashedPassword, _ := utils.HashPassword(utils.RandomString(6))

	_, user, _ := createRandomUser(hashedPassword)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testStore.CreateAccount(context.Background(), arg)
	return arg, account, err
}

func TestCreateAccount(t *testing.T) {
	arg, account, err := createRandomAccount()
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	_, accound, _ := createRandomAccount()

	fetchedAccount, err := testStore.GetAccount(context.Background(), accound.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)

	require.Equal(t, accound.Owner, fetchedAccount.Owner)
	require.Equal(t, accound.Balance, fetchedAccount.Balance)
	require.Equal(t, accound.Currency, fetchedAccount.Currency)
	require.Equal(t, accound.ID, fetchedAccount.ID)
	require.Equal(t, accound.CreatedAt, fetchedAccount.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	_, account1, _ := createRandomAccount()

	updateArgs := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}

	account2, err := testStore.UpdateAccount(context.Background(), updateArgs)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.ID, account2.ID)
	require.NotEqual(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	_, account, _ := createRandomAccount()

	err := testStore.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	deletedAccount, err := testStore.GetAccount(context.Background(), account.ID)

	require.Error(t, err)
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount()
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	fetchedAccounts, err := testStore.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, fetchedAccounts, 5)

	for _, account := range fetchedAccounts {
		require.NotEmpty(t, account)
	}
}
