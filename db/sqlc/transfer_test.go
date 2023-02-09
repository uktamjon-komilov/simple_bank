package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/uktamjon-komilov/simple_bank/util"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t);
	account2 := createRandomAccount(t);

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg);

	require.NoError(t, err);
	require.NotEmpty(t, transfer);

	require.Equal(t, transfer.FromAccountID, account1.ID);
	require.Equal(t, transfer.ToAccountID, account2.ID);
	require.Equal(t, transfer.Amount, arg.Amount);

	return transfer;
}

func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t);
}


func TestGetTransfer(t *testing.T){
	transfer1 := createRandomTransfer(t);

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID);

	require.NoError(t, err);
	require.NotEmpty(t, transfer2);

	require.Equal(t, transfer1.ID, transfer2.ID);
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID);
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID);
	require.Equal(t, transfer1.Amount, transfer2.Amount);
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second);
}

func TestListTransfers(t *testing.T){
	for i := 0; i < 4; i++{
		createRandomTransfer(t);
	}

	arg := ListTransfersParams{
		Limit: 2,
		Offset: 2,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg);

	require.NoError(t, err);
	require.Len(t, transfers, 2);

	for _, transfer := range transfers{
		require.NotEmpty(t, transfer);
	}
}