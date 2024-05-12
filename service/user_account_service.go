package service

import (
	"context"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/consts"
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
	Services    *Service
}

func NewUserAccountService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager, services *Service) *UserAccountService {
	return &UserAccountService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
		Services:    services,
	}
}

func (s UserAccountService) NewUserAccount(ctx context.Context, req *req.CreateUserAccount) (*res.CreateUserAccount, error) {
	userAccountModel := &model.UserAccount{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Document:  req.Document,
	}

	userID, err := s.RepoManager.MySQL.UserAccount.InsertUserAccount(ctx, userAccountModel)
	if err != nil {
		return nil, err
	}

	walletRes, err := s.Services.Wallet.CreateWallet(ctx, "Default wallet", userID, consts.BRL_ID)
	userAccountRes := mapper.UserAccountModelToRes(userAccountModel)

	return &res.CreateUserAccount{
		UserAccount: userAccountRes,
		Wallet:      walletRes,
	}, nil
}
