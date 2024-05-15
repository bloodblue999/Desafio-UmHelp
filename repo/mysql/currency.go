package mysql

import (
	"context"
	"github.com/bloodblue999/umhelp/model"
	"github.com/jmoiron/sqlx"
)

type Currency struct {
	cli *sqlx.DB
}

func (r *Currency) FindByCurrencyCode(ctx context.Context, code string, tx *sqlx.Tx) (*model.Currency, bool, error) {
	query := `
		SELECT 
		    currency_id,
		    code,
		    symbol,
		    created_at,
		    updated_at
		    FROM tb_currency 
         WHERE 
             code = 'BRL' AND
             deleted_at IS NULL
`
	exec := r.cli.QueryxContext
	if tx != nil {
		exec = tx.QueryxContext
	}

	rows, err := exec(ctx, query)
	if err != nil {
		return nil, false, err
	}

	defer rows.Close()

	var currencyModel model.Currency

	isFound := rows.Next()
	if !isFound {
		return nil, false, nil
	}

	if err := rows.StructScan(&currencyModel); err != nil {
		return nil, false, err
	}

	if rows.Err() != nil {
		return nil, false, rows.Err()
	}

	return &currencyModel, true, nil
}
