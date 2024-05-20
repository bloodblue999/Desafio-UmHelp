package controller

import (
	"github.com/bloodblue999/umhelp/server/controller/auth"
	"github.com/bloodblue999/umhelp/server/controller/health"
	"github.com/bloodblue999/umhelp/server/controller/useraccount"
	"github.com/bloodblue999/umhelp/server/controller/wallet"
	"github.com/bloodblue999/umhelp/service"
	"github.com/bloodblue999/umhelp/util/resutil"
	"github.com/rs/zerolog"
)

type Controller struct {
	HealthController      *health.Controller
	UserAccountController *useraccount.Controller
	WalletController      *wallet.Controller
	AuthController        *auth.Controller
}

func New(svc *service.Service, logger *zerolog.Logger) *Controller {
	resutil := resutil.New(logger)

	return &Controller{
		HealthController:      health.New(resutil),
		UserAccountController: useraccount.New(logger, resutil, svc),
		WalletController:      wallet.New(logger, resutil, svc),
		AuthController:        auth.New(logger, resutil, svc),
	}
}
