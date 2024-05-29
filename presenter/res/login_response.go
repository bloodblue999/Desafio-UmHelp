package res

import "time"

type AuthenticationTokenResponse struct {
	AccessToken    string    `json:"accessToken"`
	ExpirationDate time.Time `json:"expirationDate"`
	RefreshToken   string    `json:"refreshToken"`
}
