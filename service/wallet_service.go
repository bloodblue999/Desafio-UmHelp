package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/mapper"
	"github.com/bloodblue999/umhelp/model"
	"github.com/bloodblue999/umhelp/presenter/req"
	"github.com/bloodblue999/umhelp/presenter/res"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/jmoiron/sqlx"
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

	_, err := s.RepoManager.MySQL.Wallet.InsertWallet(ctx, walletModel, nil)
	if err != nil {
		return nil, err
	}

	return mapper.WalletModelToRes(walletModel), nil
}

func (s WalletService) NewMoneyTransaction(ctx context.Context, req *req.CreateMoneyTransaction) (*res.Wallet, error) {
	tx, err := s.RepoManager.MySQL.BeginTransaction(ctx, sql.LevelSerializable)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	senderWalletModel, err := findWalletById(ctx, s.RepoManager, tx, req.SenderID)
	if err != nil {
		return nil, err
	}

	// TODO: USE AUTHENTICATION TO VERIFY IF SENDER WALLET BELONGS TO REQUEST SENDER

	if senderWalletModel.Balance < req.MoneyValue {
		return nil, fmt.Errorf("insuficient balance in sender wallet")
	}

	receiverWalletModel, err := findWalletById(ctx, s.RepoManager, tx, req.ReceiverID)
	if err != nil {
		return nil, err
	}

	if senderWalletModel.CurrencyID != receiverWalletModel.CurrencyID {
		return nil, fmt.Errorf("diferent currency between wallets")
	}

	senderWalletModel.Balance = senderWalletModel.Balance - req.MoneyValue
	err = s.RepoManager.MySQL.Wallet.UpdateWalletBalance(ctx, senderWalletModel.ID, senderWalletModel.Balance, tx)
	if err != nil {
		return nil, err
	}

	receiverWalletModel.Balance = receiverWalletModel.Balance + req.MoneyValue
	err = s.RepoManager.MySQL.Wallet.UpdateWalletBalance(ctx, receiverWalletModel.ID, receiverWalletModel.Balance, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return mapper.WalletModelToRes(senderWalletModel), nil
}

func findWalletById(ctx context.Context, repo *repo.RepoManager, tx *sqlx.Tx, id int64) (*model.Wallet, error) {
	walletModel, found, err := repo.MySQL.Wallet.FindById(ctx, id, tx)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("wallet id `%d` not founded in database", id)
	}

	return walletModel, nil
}
