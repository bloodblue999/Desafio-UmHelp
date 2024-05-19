package mysql

import (
	"context"
	"github.com/bloodblue999/umhelp/model"
	"github.com/jmoiron/sqlx"
)

type UserAccount struct {
	cli *sqlx.DB
}

func (b *UserAccount) InsertUserAccount(ctx context.Context, userAccountModel *model.UserAccount, transaction *sqlx.Tx) (int64, error) {
	query := `INSERT INTO um_help.tb_user_account (first_name, last_name, document, password)
				VALUES (?, ?, ?, ?)`

	exec := b.cli.ExecContext
	if transaction != nil {
		exec = transaction.ExecContext
	}

	result, err := exec(ctx, query,
		userAccountModel.FirstName,
		userAccountModel.LastName,
		userAccountModel.Document,
		userAccountModel.Password,
	)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
