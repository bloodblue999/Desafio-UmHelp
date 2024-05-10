package health

import (
	"github.com/bloodblue999/umhelp/util/resutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	resutil *resutil.ResUtil
}

func New(resutil *resutil.ResUtil) *Controller {
	return &Controller{
		resutil: resutil,
	}
}

func (ctrl *Controller) HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}
