package db

import (
	"context"
	"testing"

	"github.com/prakhar30/bankoindiana/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry() (CreateEntryParams, Entry, error) {
	args := CreateEntryParams{
		AccountID: 1,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	return args, entry, err
}

func TestCreateEntry(t *testing.T) {
	args, entry, err := createRandomEntry()

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
}

func TestGetEntry(t *testing.T) {
	_, entry1, _ := createRandomEntry()

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry()
	}

	args := ListEntriesParams{
		AccountID: 1,
		Limit:     5,
		Offset:    5,
	}

	fetchedEntries, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, fetchedEntries, 5)
}
