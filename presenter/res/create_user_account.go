package res

type CreateUserAccount struct {
	UserAccount *UserAccount `json:"UserAccount"`
	Wallet      *Wallet      `json:"wallet"`
}
