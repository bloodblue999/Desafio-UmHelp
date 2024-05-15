package model

import "time"

type Currency struct {
	ID        int64      `db:"currency_id"`
	Code      string     `db:"code"`
	Symbol    string     `db:"symbol"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
