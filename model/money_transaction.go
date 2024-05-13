package model

import "time"

type MoneyTransaction struct {
	ID             int64     `db:"money_transaction_id"`
	MoneyValue     int64     `db:"money_value"`
	ProcessingData time.Time `db:"processing_data"`
	SenderID       int64     `db:"sender_id"`
	ReceiverID     int64     `db:"receiver_id"`
}
