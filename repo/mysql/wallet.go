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

	result, err := transaction.ExecContext(ctx, query,
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
