package model

import "time"

type Wallet struct {
	ID         int64      `db:"wallet_id"`
	Balance    int64      `db:"balance"`
	Alias      string     `db:"alias"`
	OwnerID    int64      `db:"owner_id"`
	CurrencyID int64      `db:"currency_id"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}
