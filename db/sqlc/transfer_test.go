package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vitaLemoTea/myBank/util"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	p := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testqueries.CreateTransfer(context.Background(), p)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, p.FromAccountID, transfer.FromAccountID)
	require.Equal(t, p.ToAccountID, transfer.ToAccountID)
	require.Equal(t, p.Amount, transfer.Amount)
	return transfer

}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	transfer2, err := testqueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

func TestGetListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}
	parm := GetListTransferParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testqueries.GetListTransfer(context.Background(), parm)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomTransfer(t)
	parm := UpdateTransferParams{
		ID:            transfer1.ID,
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer2, err := testqueries.UpdateTransfer(context.Background(), parm)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	err := testqueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	transfer2, err := testqueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, transfer2)

}
