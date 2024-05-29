package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/mapper"
	"github.com/bloodblue999/umhelp/model"
	"github.com/bloodblue999/umhelp/presenter/req"
	"github.com/bloodblue999/umhelp/presenter/res"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/golang-jwt/jwt/v5"
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

func (s WalletService) NewMoneyTransaction(ctx context.Context, req *req.CreateMoneyTransaction, claims jwt.MapClaims) (*res.Wallet, error) {
	tx, err := s.RepoManager.MySQL.BeginTransaction(ctx, sql.LevelSerializable)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	senderWalletModel, err := findWalletById(ctx, s.RepoManager, tx, req.SenderID)
	if err != nil {
		return nil, err
	}

	if senderWalletModel.Balance < req.MoneyValue {
		return nil, errors.New("insuficient balance in sender wallet")
	}

	receiverWalletModel, err := findWalletById(ctx, s.RepoManager, tx, req.ReceiverID)
	if err != nil {
		return nil, err
	}

	if senderWalletModel.CurrencyID != receiverWalletModel.CurrencyID {
		return nil, errors.New("different currency between wallets")
	}

	senderUserModel, err := findUserAccountById(ctx, s.RepoManager, tx, senderWalletModel.OwnerID)
	if err != nil {
		return nil, err
	}

	subjectID, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	if senderUserModel.PublicID != subjectID {
		return nil, errors.New("sender owner id is different from authentication id")
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

func findUserAccountById(ctx context.Context, repo *repo.RepoManager, tx *sqlx.Tx, id int64) (*model.UserAccount, error) {
	userModel, isFound, err := repo.MySQL.UserAccount.SelectUserByID(ctx, id, tx)
	if err != nil {
		return nil, err
	}

	if !isFound {
		return nil, fmt.Errorf("owner user id `%d` not founded in database", id)
	}

	return userModel, nil
}
