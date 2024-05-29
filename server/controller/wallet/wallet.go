package wallet

import (
	"github.com/bloodblue999/umhelp/consts"
	"github.com/bloodblue999/umhelp/service"
	"github.com/bloodblue999/umhelp/util/claimsutil"
	"github.com/bloodblue999/umhelp/util/resutil"
	"github.com/bloodblue999/umhelp/validation"
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

	claims, err := claimsutil.ParseToMapClaims(ctx.Get(consts.ClaimsName))
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	data, err := ctrl.services.Wallet.NewMoneyTransaction(ctx.Request().Context(), req, claims)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(data, nil, http.StatusCreated))
}
