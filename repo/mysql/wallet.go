package mysql

import (
	"context"
	"github.com/bloodblue999/umhelp/model"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	cli *sqlx.DB
}

func (r Wallet) InsertWallet(ctx context.Context, walletModel *model.Wallet) error {
	query := `INSERT INTO tb_wallet (alias, owner_id, currency_id) 
		VALUES (?,?,?)`

	_, err := r.cli.ExecContext(ctx, query,
		walletModel.Alias,
		walletModel.OwnerID,
		walletModel.CurrencyID,
	)
	if err != nil {
		return err
	}

	return nil
}
