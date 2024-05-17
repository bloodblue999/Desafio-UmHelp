package req

type LoginRequest struct {
	Document string `json:"document"`
	Password string `json:"password"`
}
