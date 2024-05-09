package model

import "time"

type UserAccount struct {
	ID        int64      `db:"id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Document  string     `db:"document"`
	Balance   int64      `db:"balance"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
