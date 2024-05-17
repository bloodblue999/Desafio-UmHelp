package router

import (
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/server/controller"
	"github.com/labstack/echo/v4"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {

	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	userAccount := root.Group("/useraccount")
	userAccount.POST("", ctrl.UserAccountController.HandleNewUserAccount)

	wallet := root.Group("/wallet")
	wallet.POST("/transaction", ctrl.WalletController.HandleNewMoneyTransaction)

	auth := root.Group("/auth")
	auth.POST("/login", ctrl.AuthController.HandleLoginRequest)

}
