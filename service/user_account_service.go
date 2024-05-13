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

func NewUserAccountService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *UserAccountService {
	return &UserAccountService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s UserAccountService) NewUserAccount(ctx context.Context, req *req.CreateUserAccount) (*res.CreateUserAccount, error) {
	userAccountModel := &model.UserAccount{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Document:  req.Document,
	}

	transaction, err := s.RepoManager.MySQL.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer transaction.Rollback()

	userID, err := s.RepoManager.MySQL.UserAccount.InsertUserAccount(ctx, userAccountModel, transaction)
	if err != nil {
		return nil, err
	}

	walletModel := &model.Wallet{
		Alias:      "Default wallet",
		OwnerID:    userID,
		CurrencyID: consts.BrlId,
	}
	_, err = s.RepoManager.MySQL.Wallet.InsertWallet(ctx, walletModel, transaction)
	if err != nil {
		return nil, err
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	userAccountRes := mapper.UserAccountModelToRes(userAccountModel)
	walletRes := mapper.WalletModelToRes(walletModel)

	return &res.CreateUserAccount{
		UserAccount: userAccountRes,
		Wallet:      walletRes,
	}, nil
}
