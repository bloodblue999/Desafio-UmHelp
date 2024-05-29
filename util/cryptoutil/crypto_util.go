package cryptoutil

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/consts"
	"github.com/bloodblue999/umhelp/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type CryptoUtil struct {
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
	cfg        *config.Config
}

func NewCryptUtil(cfg *config.Config) (*CryptoUtil, error) {
	privateKey, err := parsePrivateKey(cfg.CryptoConfig.JWSPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey, err := parsePublicKey(cfg.CryptoConfig.JWSPublicKey)
	if err != nil {
		return nil, err
	}

	return &CryptoUtil{
		publicKey:  publicKey,
		privateKey: privateKey,
		cfg:        cfg,
	}, nil

}

func parsePrivateKey(key string) (ed25519.PrivateKey, error) {
	// TODO: CHANGE PRIVATE KEY PARSE
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	if len(keyBytes) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("private Jwt Key Size is invalid")
	}

	return keyBytes, nil
}

func parsePublicKey(key string) (ed25519.PublicKey, error) {
	// TODO: CHANGE PUBLIC KEY PARSE

	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	if len(keyBytes) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("public Jwt Key Size is invalid")
	}

	return keyBytes, nil
}

type TokenResult struct {
	SignID         string
	AccessToken    string
	ExpirationTime time.Time
	RefreshToken   string
}

func (crypto *CryptoUtil) SignUserToken(cfg *config.Config, subjectPublicID string) (*TokenResult, error) {
	expirationTime := time.Now().Add(time.Hour * time.Duration(crypto.cfg.CryptoConfig.JWSExpirationTimeInHours))
	signID := uuid.New().String()

	jwtAccessTokenConfig := jwt.NewWithClaims(
		jwt.SigningMethodEdDSA,
		entity.TokenJWT{
			SignID:     signID,
			Issuer:     cfg.InternalConfig.ServiceName,
			Subject:    subjectPublicID,
			IssuedAt:   time.Now().Unix(),
			Expiration: expirationTime.Unix(),
			TokenType:  consts.AccessTokenType,
		},
	)
	accessTokenStr, err := jwtAccessTokenConfig.SignedString(crypto.privateKey)
	if err != nil {
		return nil, err
	}

	jwtRefreshTokenConfig := jwt.NewWithClaims(
		jwt.SigningMethodEdDSA,
		entity.TokenJWT{
			SignID:    signID,
			Issuer:    cfg.InternalConfig.ServiceName,
			Subject:   subjectPublicID,
			IssuedAt:  time.Now().Unix(),
			TokenType: consts.RefreshTokenType,
		},
	)
	refreshTokenStr, err := jwtRefreshTokenConfig.SignedString(crypto.privateKey)
	if err != nil {
		return nil, err
	}

	return &TokenResult{
		SignID:         signID,
		AccessToken:    accessTokenStr,
		ExpirationTime: expirationTime,
		RefreshToken:   refreshTokenStr,
	}, nil
}

func (crypto *CryptoUtil) VerifyASignatureToken(token string) (jwt.Claims, error) {
	var keyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return crypto.publicKey, nil }
	validator, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return nil, err
	}

	return validator.Claims, err
}

func (crypto *CryptoUtil) HashString(password string) string {
	mac := hmac.New(sha256.New, []byte(crypto.cfg.CryptoConfig.HS256Password))
	mac.Write([]byte(password))

	return hex.EncodeToString(mac.Sum(nil))
}
