package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/myusufirfanh/go-demo-service/domain/addon"
	"github.com/myusufirfanh/go-demo-service/shared/util"
)

type handlerAddon struct {
	usecase addon.Usecase
}

func AddAddonHandler(e *echo.Echo, usecase addon.Usecase) {
	handler := handlerAddon{
		usecase: usecase,
	}

	e.GET("/api/addon/get", handler.GetAddon)
}

func (h handlerAddon) GetAddon(c echo.Context) error {
	ac := c.(*util.CustomApplicationContext)

	resp, err := h.usecase.GetAddon(c, 1)
	if err != nil {
		return ac.CustomResponse("failed", resp, "Failed to Trigger Payment Reminder!", http.StatusInternalServerError, nil, "")
	}
	return ac.CustomResponse("success", resp, "Success to Trigger Payment Reminder!", http.StatusOK, nil, "")
}
