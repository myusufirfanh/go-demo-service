package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/myusufirfanh/go-demo-service/domain/product"
	"github.com/myusufirfanh/go-demo-service/shared/util"
)

type handlerProduct struct {
	usecase product.Usecase
}

func AddProductHandler(e *echo.Echo, usecase product.Usecase) {
	handler := handlerProduct{
		usecase: usecase,
	}

	e.GET("/api/product/get", handler.GetProduct)
}

func (h handlerProduct) GetProduct(c echo.Context) error {
	ac := c.(*util.CustomApplicationContext)

	resp, err := h.usecase.GetProduct(c, 1)
	if err != nil {
		return ac.CustomResponse("failed", resp, "Failed to Trigger Payment Reminder!", http.StatusInternalServerError, nil, "")
	}
	return ac.CustomResponse("success", resp, "Success to Trigger Payment Reminder!", http.StatusOK, nil, "")
}
