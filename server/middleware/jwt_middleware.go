package middleware

import (
	"github.com/bloodblue999/umhelp/consts"
	"github.com/bloodblue999/umhelp/util/cryptoutil"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTTokenAuthentication(cryptoUtil *cryptoutil.CryptoUtil, requestedTokenType string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

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

			if authorizationName != consts.AuthenticationTypeString {
				return ctx.JSON(http.StatusUnauthorized, "invalid authentication method")
			}

			claims, err := cryptoUtil.VerifyASignatureToken(token)
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, err.Error())
			}

			claimsMap := claims.(jwt.MapClaims)

			tokenType := claimsMap[consts.TokenTypeParamether]

			if tokenType != requestedTokenType {
				return ctx.JSON(http.StatusUnauthorized, "invalid token type")
			}

			ctx.Set(consts.ClaimsName, claimsMap)
			return next(ctx)
		}

	}
}
