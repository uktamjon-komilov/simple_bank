package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/uktamjon-komilov/simple_bank/util"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t);
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg);

	require.NoError(t, err);
	require.NotEmpty(t, entry);

	require.Equal(t, entry.AccountID, account.ID);
	require.Equal(t, entry.Amount, arg.Amount);

	return entry;
}

func TestCreateEntry(t *testing.T){
	createRandomEntry(t);
}

func TestGetEntry(t *testing.T){
	entry1 := createRandomEntry(t);

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID);

	require.NoError(t, err);
	require.NotEmpty(t, entry2);

	require.Equal(t, entry1.ID, entry2.ID);
	require.Equal(t, entry1.AccountID, entry2.AccountID);
	require.Equal(t, entry1.Amount, entry2.Amount);
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second);
}

func TestListEntries(t *testing.T){
	for i := 0; i < 4; i++ {
		createRandomEntry(t);
	}

	arg := ListEntriesParams{
		Limit: 2,
		Offset: 2,
	}
	
	entries, err := testQueries.ListEntries(context.Background(), arg);

	require.NoError(t, err);
	require.Len(t, entries, 2);

	for _, entry := range entries {
		require.NotEmpty(t, entry);
	}
}