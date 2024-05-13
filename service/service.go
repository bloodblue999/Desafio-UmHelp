package service

import (
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/rs/zerolog"
)

type Service struct {
	UserAccount *UserAccountService
	Wallet      *WalletService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	return &Service{
		Wallet:      NewWalletService(cfg, logger, repo),
		UserAccount: NewUserAccountService(cfg, logger, repo),
	}
}
