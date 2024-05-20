package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/presenter/req"
	"github.com/bloodblue999/umhelp/presenter/res"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/bloodblue999/umhelp/util/cryptoutil"
	"github.com/rs/zerolog"
	"time"
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

func (s AuthService) Login(ctx context.Context, req *req.LoginRequest) (*res.LoginRequest, error) {
	userModel, isFound, err := s.RepoManager.MySQL.UserAccount.SelectUserAccountByDocument(ctx, req.Document, nil)
	if err != nil {
		return nil, err
	}

	if !isFound {
		return nil, fmt.Errorf("cannot found document %s in database", req.Document)
	}

	hashedReqPassword := s.CryptoUtil.HashPassword(req.Password)

	if hashedReqPassword != userModel.Password {
		return nil, errors.New("login error, invalid password")
	}

	token, err := s.CryptoUtil.CreateASignedToken(userModel.PublicID)
	if err != nil {
		return nil, errors.New("error to generate token")
	}

	expirationDate := time.Now().Add(time.Hour * time.Duration(s.Config.CryptoConfig.JWSExpirationTimeInHours))

	return &res.LoginRequest{
		Token:          token,
		ExpirationDate: expirationDate,
	}, nil
}
