package middleware

import (
	"github.com/bloodblue999/umhelp/util/cryptoutil"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Authenticate struct {
	cryptoUtil cryptoutil.CryptoUtil
}

func NewAuthenticateMiddleWare(cryptoUtil *cryptoutil.CryptoUtil) *Authenticate {
	return &Authenticate{
		cryptoUtil: *cryptoUtil,
	}
}

func (m *Authenticate) AuthenticatedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorizationHeader := ctx.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return ctx.JSON(http.StatusUnauthorized, "invalid authentication, no authentication detected")
		}

		authenticationSplited := strings.Split(authorizationHeader, " ")
		if len(authenticationSplited) != 2 {
			return ctx.JSON(http.StatusUnauthorized, "invalid authentication")
		}

		authorizationName, token := authenticationSplited[0], authenticationSplited[1]

		if authorizationName != "Bearer" {
			return ctx.JSON(http.StatusUnauthorized, "invalid authentication method")
		}

		claims, err := m.cryptoUtil.VerifyASignatureToken(token)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, err.Error())
		}

		ctx.Set("claims", claims)
		return next(ctx)
	}
}
