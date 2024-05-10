package service

import (
	"context"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/model"
	"github.com/bloodblue999/umhelp/presenter/req"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/rs/zerolog"
)

type UserAccountService struct {
	Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
}

func NewUserAccountService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *UserAccountService {
	return &UserAccountService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s UserAccountService) NewUserAccount(ctx context.Context, req *req.CreateUserAccount) error {

	userAccountModel := &model.UserAccount{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Document:  req.Document,
	}

	return s.RepoManager.MySQL.UserAccount.InsertUserAccount(ctx, userAccountModel)
}
