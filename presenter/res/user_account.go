package res

type UserAccount struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Document  string `json:"document"`
	Balance   int64  `json:"balance"`
}
