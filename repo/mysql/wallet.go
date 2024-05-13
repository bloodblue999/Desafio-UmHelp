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
