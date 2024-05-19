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

func (b *UserAccount) SelectUserAccountByDocument(ctx context.Context, document string, transaction *sqlx.Tx) (*model.UserAccount, bool, error) {
	query := `SELECT *
		FROM tb_user_account
		WHERE document = ? AND 
		      deleted_at IS NULL`

	exec := b.cli.QueryxContext
	if transaction != nil {
		exec = transaction.QueryxContext
	}

	rows, err := exec(ctx, query, document)
	if err != nil {
		return nil, false, err
	}

	defer rows.Close()

	var userAccountModel model.UserAccount

	isFound := rows.Next()
	if !isFound {
		return nil, false, nil
	}

	if err := rows.StructScan(&userAccountModel); err != nil {
		return nil, false, err
	}

	if rows.Err() != nil {
		return nil, false, rows.Err()
	}

	return &userAccountModel, true, nil
}
