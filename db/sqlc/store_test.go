package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 10
	amount := int64(10)

	errs := make(chan error)

	for i := 0; i < n; i++ {

		From_Account := account1
		To_Account := account2

		if i % 2 == 1 {
			From_Account = account2
			To_Account = account1
		}
		
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: From_Account.ID,
				ToAccountID:   To_Account.ID,
				Amount:        amount,
			})
			errs <- err
		}()
		

	}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	//check updated account

	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}

