package res

type Wallet struct {
	Alias      string `json:"alias"`
	Balance    int64  `json:"balance"`
	CurrencyID int64  `json:"currencyID"`
}
