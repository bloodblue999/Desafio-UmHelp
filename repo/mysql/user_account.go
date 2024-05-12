package mysql

import (
	"context"
	"github.com/bloodblue999/umhelp/model"
	"github.com/jmoiron/sqlx"
)

type UserAccount struct {
	cli *sqlx.DB
}

func (b *UserAccount) InsertUserAccount(ctx context.Context, userAccountModel *model.UserAccount) (int64, error) {
	query := `INSERT INTO um_help.tb_user_account (first_name, last_name, document)
				VALUES (?, ?, ?)`

	result, err := b.cli.ExecContext(ctx, query,
		userAccountModel.FirstName,
		userAccountModel.LastName,
		userAccountModel.Document,
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
