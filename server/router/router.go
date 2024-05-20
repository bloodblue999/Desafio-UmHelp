package router

import (
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/server/controller"
	"github.com/bloodblue999/umhelp/server/middleware"
	"github.com/bloodblue999/umhelp/util/cryptoutil"
	"github.com/labstack/echo/v4"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller, cryptoUtil *cryptoutil.CryptoUtil) {
	authenticatedMiddleWare := middleware.NewAuthenticateMiddleWare(cryptoUtil)

	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	userAccount := root.Group("/useraccount")
	userAccount.POST("", ctrl.UserAccountController.HandleNewUserAccount)

	wallet := root.Group("/wallet", authenticatedMiddleWare.AuthenticatedMiddleware)
	wallet.POST("/transaction", ctrl.WalletController.HandleNewMoneyTransaction, authenticatedMiddleWare.AuthenticatedMiddleware)

	auth := root.Group("/auth")
	auth.POST("/login", ctrl.AuthController.HandleLoginRequest)

}
