package service

import (
	"context"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/mapper"
	"github.com/bloodblue999/umhelp/model"
	"github.com/bloodblue999/umhelp/presenter/res"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/rs/zerolog"
)

type WalletService struct {
	Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
}

func NewWalletService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *WalletService {
	return &WalletService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s WalletService) CreateWallet(ctx context.Context, alias string, ownerID, currencyID int64) (*res.Wallet, error) {
	walletModel := &model.Wallet{
		Alias:      alias,
		OwnerID:    ownerID,
		CurrencyID: currencyID,
	}

	transaction, err := s.RepoManager.MySQL.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer transaction.Rollback()

	_, err = s.RepoManager.MySQL.Wallet.InsertWallet(ctx, walletModel, transaction)
	if err != nil {
		return nil, err
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	return mapper.WalletModelToRes(walletModel), nil
}
