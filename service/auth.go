package service

import (
	"context"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/presenter/req"
	"github.com/bloodblue999/umhelp/presenter/res"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/rs/zerolog"
)

type AuthService struct {
	Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
	Services    *Service
}

func NewAuthService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *AuthService {
	return &AuthService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s AuthService) Login(ctx context.Context, req *req.LoginRequest) (*res.LoginRequest, error) {
	return nil, nil
}
