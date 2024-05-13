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
	service := &Service{}
	service.Wallet = NewWalletService(cfg, logger, repo)
	service.UserAccount = NewUserAccountService(cfg, logger, repo, service)
	return service
}
