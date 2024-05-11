package useraccount

import (
	"github.com/bloodblue999/umhelp/service"
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

func (ctrl *Controller) HandleNewUserAccount(ctx echo.Context) error {
	req, err := validation.GetAndValidateCreateUserAccount(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	data, err := ctrl.services.UserAccount.NewUserAccount(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(data, nil, http.StatusCreated))
}
