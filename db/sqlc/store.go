package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	Querier
}

type DBStore struct {
	*Queries
	conn *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &DBStore{
		Queries: New(connPool),
		conn:    connPool,
	}
}

func (store *DBStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return fmt.Errorf("transaction error: %v", err)
	}

	return tx.Commit(ctx)
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

func (store *DBStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
			if err != nil {
				return err
			}
			result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      arg.FromAccountID,
				Balance: account1.Balance - arg.Amount,
			})
			if err != nil {
				return err
			}

			account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
			if err != nil {
				return err
			}
			result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      arg.ToAccountID,
				Balance: account2.Balance + arg.Amount,
			})
			if err != nil {
				return err
			}
		} else {
			account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
			if err != nil {
				return err
			}
			result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      arg.ToAccountID,
				Balance: account2.Balance + arg.Amount,
			})
			if err != nil {
				return err
			}

			account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
			if err != nil {
				return err
			}
			result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      arg.FromAccountID,
				Balance: account1.Balance - arg.Amount,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}
