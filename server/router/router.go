package router

import (
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/consts"
	"github.com/bloodblue999/umhelp/server/controller"
	"github.com/bloodblue999/umhelp/server/middleware"
	"github.com/bloodblue999/umhelp/util/cryptoutil"
	"github.com/labstack/echo/v4"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller, cryptoUtil *cryptoutil.CryptoUtil) {
	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	userAccount := root.Group("/useraccount")
	userAccount.POST("", ctrl.UserAccountController.HandleNewUserAccount)

	wallet := root.Group("/wallet")
	wallet.POST("/transaction",
		ctrl.WalletController.HandleNewMoneyTransaction,
		middleware.JWTTokenAuthentication(cryptoUtil, consts.AccessTokenType),
	)

	auth := root.Group("/auth")
	auth.POST("/login", ctrl.AuthController.HandleLoginRequest)
	auth.GET("/refresh",
		ctrl.AuthController.HandleRefreshTokenRequest,
		middleware.JWTTokenAuthentication(cryptoUtil, consts.RefreshTokenType),
	)

}
