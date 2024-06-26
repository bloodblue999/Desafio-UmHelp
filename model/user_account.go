package model

import "time"

type UserAccount struct {
	ID        int64      `db:"user_account_id"`
	PublicID  string     `db:"public_id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Document  string     `db:"document"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
