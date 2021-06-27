package addon

import (
	"github.com/labstack/echo"
	"github.com/myusufirfanh/go-demo-service/models"
)

type Usecase interface {
	GetAddon(c echo.Context, addonID int64) (*models.Addon, error)
	GetAllAddonsByProductID(c echo.Context, productID int64) ([]models.Addon, error)
}

type Repository interface {
	GetAddon(addonID int64) (*models.Addon, error)
	GetAllAddonsByProductID(productID int64) ([]models.Addon, error)
}
