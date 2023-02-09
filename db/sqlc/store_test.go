package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB);

	account1 := createRandomAccount(t);
	account2 := createRandomAccount(t);

	n := 10;
	amount:=int64(15);

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i:= 0; i < n; i++{
		go func(){
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID: account2.ID,
				Amount: amount,
			})

			errs <- err;
			results <- result;
		}()
	}


	for i:=0; i< n; i++ {
		err := <-errs;
		require.NoError(t, err);
		
		result := <-results;
		require.NotEmpty(t, result);

		// check transfer
		transer := result.Transfer
		require.NotEmpty(t, transer);
		require.Equal(t, account1.ID, transer.FromAccountID);
		require.Equal(t, account2.ID, transer.ToAccountID);
		require.Equal(t, amount, transer.Amount);
		require.NotZero(t, transer.ID);
		require.NotZero(t, transer.CreatedAt);

		_, err = store.GetTransfer(context.Background(), transer.ID);
		require.NoError(t, err);

		// check from entry
		fromEntry := result.FromEntry;
		require.NotEmpty(t, fromEntry);
		require.Equal(t, account1.ID, fromEntry.AccountID);
		require.Equal(t, -amount, fromEntry.Amount);
		require.NotZero(t, fromEntry.ID);
		require.NotZero(t, fromEntry.CreatedAt);

		_, err = store.GetEntry(context.Background(), fromEntry.ID);
		require.NoError(t, err);

		// check to entry
		toEntry := result.ToEntry;
		require.NotEmpty(t, toEntry);
		require.Equal(t, account2.ID, toEntry.AccountID);
		require.Equal(t, amount, toEntry.Amount);
		require.NotZero(t, toEntry.ID);
		require.NotZero(t, toEntry.CreatedAt);

		_, err = store.GetEntry(context.Background(), toEntry.ID);
		require.NoError(t, err);

		// TODO: check accounts' balance
	}
}