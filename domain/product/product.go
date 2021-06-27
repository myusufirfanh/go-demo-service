package product

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/myusufirfanh/go-demo-service/models"
)

type Usecase interface {
	GetProduct(c echo.Context, productID int64) (*models.Product, error)
	GetAddonsByProductID(c echo.Context, productID int64) ([]models.Addon, error)
}

type Repository interface {
	GetProduct(c echo.Context, productID int64) (*models.Product, error)
	TxInsertProduct(c echo.Context, tx *gorm.DB, data *models.Product) (*models.Product, error)
}
