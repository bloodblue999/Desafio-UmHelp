package claimsutil

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func ParseToMapClaims(claimsInterface interface{}) (jwt.MapClaims, error) {
	claims, ok := claimsInterface.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error, cannot convert claims")
	}

	return claims, nil
}
