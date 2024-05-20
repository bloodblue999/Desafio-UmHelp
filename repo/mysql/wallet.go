package mysql

import (
	"context"
	"github.com/bloodblue999/umhelp/model"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	cli *sqlx.DB
}

func (r Wallet) InsertWallet(ctx context.Context, walletModel *model.Wallet, transaction *sqlx.Tx) (int64, error) {
	query := `INSERT INTO tb_wallet (alias, owner_id, currency_id) 
		VALUES (?,?,?)`

	exec := r.cli.ExecContext
	if transaction != nil {
		exec = transaction.ExecContext
	}

	result, err := exec(ctx, query,
		walletModel.Alias,
		walletModel.OwnerID,
		walletModel.CurrencyID,
	)
	if err != nil {
		return 0, err
	}

	walletId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return walletId, nil
}

func (r *Wallet) FindById(ctx context.Context, id int64, tx *sqlx.Tx) (*model.Wallet, bool, error) {
	query := `
		SELECT 
		    wallet_id,
		    balance,
		    alias,
		    owner_id,
		    currency_id,
		    created_at,
		    updated_at
		    FROM tb_wallet
    	     WHERE
    	         wallet_id = ? AND 
    	         deleted_at IS NULL
    	         `

	exec := r.cli.QueryxContext
	if tx != nil {
		exec = tx.QueryxContext
	}

	rows, err := exec(ctx, query, id)
	if err != nil {
		return nil, false, err
	}

	defer rows.Close()

	var walletModel model.Wallet

	isFound := rows.Next()
	if !isFound {
		return nil, false, nil
	}

	if err := rows.StructScan(&walletModel); err != nil {
		return nil, false, err
	}

	if rows.Err() != nil {
		return nil, false, rows.Err()
	}

	return &walletModel, true, nil
}

func (r *Wallet) UpdateWalletBalance(ctx context.Context, id int64, balance int64, tx *sqlx.Tx) error {
	query := `
		UPDATE tb_wallet
		SET balance = ?
		WHERE wallet_id = ?
		`
	exec := r.cli.ExecContext
	if tx != nil {
		exec = tx.ExecContext
	}

	_, err := exec(ctx, query, balance, id)
	if err != nil {
		return err
	}

	return nil
}
