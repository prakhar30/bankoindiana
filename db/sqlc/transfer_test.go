package db

import (
	"context"
	"testing"

	"github.com/prakhar30/bankoindiana/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer() (CreateTransferParams, Transfer, error) {
	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testStore.CreateTransfer(context.Background(), arg)
	return arg, transfer, err
}

func TestCreateTransfer(t *testing.T) {
	arg, transfer, _ := CreateRandomTransfer()

	fetchedTransfer, err := testStore.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedTransfer)

	require.Equal(t, arg.FromAccountID, fetchedTransfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, fetchedTransfer.ToAccountID)
	require.Equal(t, arg.Amount, fetchedTransfer.Amount)
}

func TestGetTransfer(t *testing.T) {
	_, transfer, _ := CreateRandomTransfer()

	fetchedTransfer, err := testStore.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedTransfer)

	require.Equal(t, transfer.ID, fetchedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, fetchedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, fetchedTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, fetchedTransfer.Amount)
	require.Equal(t, transfer.CreatedAt, fetchedTransfer.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer()
	}

	args := ListTransfersParams{
		FromAccountID: 1,
		ToAccountID:   2,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testStore.ListTransfers(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
