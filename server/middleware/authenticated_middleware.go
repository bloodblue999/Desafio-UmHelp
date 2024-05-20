package middleware

import (
	"fmt"
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
		authorizationName, token := authenticationSplited[0], authenticationSplited[1]

		if authorizationName != "Bearer" {
			return ctx.JSON(http.StatusUnauthorized, "invalid authentication method")
		}

		claims, err := m.cryptoUtil.VerifyASignatureToken(token)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, err.Error())
		}

		subjectId, err := claims.GetSubject()
		if err != nil {
			fmt.Println(claims)
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.Set("subjectID", subjectId)
		return next(ctx)
	}
}
