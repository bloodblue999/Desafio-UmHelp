package service

import (
	"context"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/mapper"
	"github.com/bloodblue999/umhelp/model"
	"github.com/bloodblue999/umhelp/presenter/req"
	"github.com/bloodblue999/umhelp/presenter/res"
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

func (s UserAccountService) NewUserAccount(ctx context.Context, req *req.CreateUserAccount) (*res.UserAccount, error) {

	userAccountModel := &model.UserAccount{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Document:  req.Document,
		Balance:   0,
	}

	err := s.RepoManager.MySQL.UserAccount.InsertUserAccount(ctx, userAccountModel)
	if err != nil {
		return nil, err
	}

	return mapper.UserAccountModelToRes(userAccountModel), nil
}
