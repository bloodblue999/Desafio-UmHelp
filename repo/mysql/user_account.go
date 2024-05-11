package mysql

import (
	"context"
	"github.com/bloodblue999/umhelp/model"
	"github.com/jmoiron/sqlx"
)

type UserAccount struct {
	cli *sqlx.DB
}

func (b *UserAccount) InsertUserAccount(ctx context.Context, userAccountModel *model.UserAccount) error {
	query := `INSERT INTO um_help.tb_user_account (first_name, last_name, document, balance)
				VALUES (?, ?, ?, ?)`

	_, err := b.cli.ExecContext(ctx, query,
		userAccountModel.FirstName,
		userAccountModel.LastName,
		userAccountModel.Document,
		userAccountModel.Balance,
	)
	if err != nil {
		return err
	}

	return nil
}
