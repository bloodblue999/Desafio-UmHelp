package wallet

import (
	"errors"
	"github.com/bloodblue999/umhelp/service"
	"github.com/bloodblue999/umhelp/util/resutil"
	"github.com/bloodblue999/umhelp/validation"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
)

type Controller struct {
	logger   *zerolog.Logger
	resutil  *resutil.ResUtil
	services *service.Service
}

func New(logger *zerolog.Logger, resutil *resutil.ResUtil, services *service.Service) *Controller {
	return &Controller{
		logger:   logger,
		resutil:  resutil,
		services: services,
	}
}

func (ctrl *Controller) HandleNewMoneyTransaction(ctx echo.Context) error {
	req, err := validation.GetAndValidateMoneyTransaction(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}
	claims, err := getClaims(ctx.Get("claims"))
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	data, err := ctrl.services.Wallet.NewMoneyTransaction(ctx.Request().Context(), req, claims)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(data, nil, http.StatusCreated))
}

func getClaims(claimsInterface interface{}) (*jwt.MapClaims, error) {
	claims, ok := claimsInterface.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error, cannot convert claims")
	}

	return &claims, nil
}
