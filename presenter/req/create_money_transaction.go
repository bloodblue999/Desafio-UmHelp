package req

type CreateMoneyTransaction struct {
	MoneyValue int64 `json:"moneyValue"`
	SenderID   int64 `json:"senderID"`
	ReceiverID int64 `json:"receiverID"`
}
