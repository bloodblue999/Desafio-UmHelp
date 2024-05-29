package entity

import "github.com/golang-jwt/jwt/v5"

type TokenJWT struct {
	SignID     string `json:"jti"`
	Issuer     string `json:"iss"`
	Subject    string `json:"sub"`
	IssuedAt   int64  `json:"iat"`
	Expiration int64  `json:"exp"`
	TokenType  string `json:"type"`
	jwt.RegisteredClaims
}
