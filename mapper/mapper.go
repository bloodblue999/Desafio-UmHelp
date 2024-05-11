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
		Balance:   userAccountModel.Balance,
	}
}
