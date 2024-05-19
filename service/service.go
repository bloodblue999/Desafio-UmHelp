package service

import (
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/bloodblue999/umhelp/util/resutil"
	"github.com/rs/zerolog"
)

type Service struct {
	UserAccount *UserAccountService
	Wallet      *WalletService
	AuthService *AuthService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	cryptoUtl, err := resutil.NewCryptService(cfg)
	if err != nil {
		panic(err)
	}

	return &Service{
		Wallet:      NewWalletService(cfg, logger, repo),
		UserAccount: NewUserAccountService(cfg, logger, repo, cryptoUtl),
		AuthService: NewAuthService(cfg, logger, repo, cryptoUtl),
	}
}
