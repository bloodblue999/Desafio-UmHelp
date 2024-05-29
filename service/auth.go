package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/consts"
	"github.com/bloodblue999/umhelp/presenter/req"
	"github.com/bloodblue999/umhelp/presenter/res"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/bloodblue999/umhelp/util/cryptoutil"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
)

type AuthService struct {
	Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
	CryptoUtil  *cryptoutil.CryptoUtil
}

func NewAuthService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager, cryptoUtil *cryptoutil.CryptoUtil) *AuthService {
	return &AuthService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
		CryptoUtil:  cryptoUtil,
	}
}

func (s AuthService) Login(ctx context.Context, req *req.LoginRequest) (*res.AuthenticationTokenResponse, error) {
	userModel, isFound, err := s.RepoManager.MySQL.UserAccount.SelectUserAccountByDocument(ctx, req.Document, nil)
	if err != nil {
		return nil, err
	}

	if !isFound {
		return nil, fmt.Errorf("cannot found document %s in database", req.Document)
	}

	hashedReqPassword := s.CryptoUtil.HashString(req.Password)

	if hashedReqPassword != userModel.Password {
		return nil, errors.New("login error, invalid password")
	}

	result, err := s.CryptoUtil.SignUserToken(s.Config, userModel.PublicID)
	if err != nil {
		return nil, errors.New("error to generate token")
	}

	if err := s.RepoManager.Redis.Util.SetStructure(ctx, result.SignID, result, 0); err != nil {
		return nil, err
	}

	return &res.AuthenticationTokenResponse{
		AccessToken:    result.AccessToken,
		ExpirationDate: result.ExpirationTime,
		RefreshToken:   result.RefreshToken,
	}, nil
}

func (s AuthService) Refresh(ctx context.Context, claims jwt.MapClaims) (*res.AuthenticationTokenResponse, error) {
	signID := claims[consts.SignatureIDParamether].(string)

	_, err := s.RepoManager.Redis.Util.GetString(ctx, signID)
	if err != nil {
		return nil, err
	}

	userID, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	if err := s.RepoManager.Redis.Util.DelString(ctx, signID); err != nil {
		return nil, err
	}

	result, err := s.CryptoUtil.SignUserToken(s.Config, userID)
	if err != nil {
		return nil, err
	}

	if err := s.RepoManager.Redis.Util.SetStructure(ctx, result.SignID, result, 0); err != nil {
		return nil, err
	}

	return &res.AuthenticationTokenResponse{
		AccessToken:    result.AccessToken,
		ExpirationDate: result.ExpirationTime,
		RefreshToken:   result.RefreshToken,
	}, nil
}
