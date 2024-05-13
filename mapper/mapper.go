package mapper

import (
	"github.com/bloodblue999/umhelp/model"
	"github.com/bloodblue999/umhelp/presenter/res"
)

func UserAccountModelToRes(userAccountModel *model.UserAccount) *res.UserAccount {
	if userAccountModel == nil {
		return nil
	}

	return &res.UserAccount{
		FirstName: userAccountModel.FirstName,
		LastName:  userAccountModel.LastName,
		Document:  userAccountModel.Document,
	}
}

func WalletModelToRes(walletModel *model.Wallet) *res.Wallet {
	if walletModel == nil {
		return nil
	}

	return &res.Wallet{
		Alias:      walletModel.Alias,
		Balance:    walletModel.Balance,
		CurrencyID: walletModel.CurrencyID,
	}
}
