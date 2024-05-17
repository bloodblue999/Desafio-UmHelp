package res

import "time"

type LoginRequest struct {
	Token          string    `json:"token"`
	ExpirationDate time.Time `json:"expirationDate"`
}
