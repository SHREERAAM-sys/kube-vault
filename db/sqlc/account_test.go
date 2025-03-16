package db

import (
	"context"
	"database/sql"
	"github.com/SHREERAAM-sys/kube-vault/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Accounts {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)      //check weather the error is nil, or else automatically fail the test
	require.NotEmpty(t, account) //check weather the account obj is not empty

	//testing the input args and the created account obj created in the database
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	//check if pgsql automatically create value for the belo by checking not zero
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {

	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {

	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.WithinDurationf(t, account1.CreatedAt, account2.CreatedAt, time.Second, "created at sendos should match")
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance) //arr the balance is the current balance
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDurationf(t, account1.CreatedAt, account2.CreatedAt, time.Second, "created at sendos should match")

}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	//to ensure the account is deleted, trying to get the deleted account obj for the database
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)                //check if there is err, error should be there since we are retrieving obj of the deleted account
	require.Error(t, err, sql.ErrNoRows) //for precise checking
	require.Empty(t, account2)

}

func TestListAccounts(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	} //will get at least 5 accounts

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
