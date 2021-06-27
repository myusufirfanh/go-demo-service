package usecase

import (
	"github.com/labstack/echo"
	"github.com/myusufirfanh/go-demo-service/domain/addon"
	"github.com/myusufirfanh/go-demo-service/domain/product"
	"github.com/myusufirfanh/go-demo-service/models"
	"github.com/myusufirfanh/go-demo-service/shared/util"
)

type usecase struct {
	repository   product.Repository
	addonUsecase addon.Usecase
}

func NewProductUsecase(repository product.Repository, addonUsecase addon.Usecase) product.Usecase {
	return &usecase{
		repository:   repository,
		addonUsecase: addonUsecase,
	}
}

// Get a product from DB
func (u usecase) GetProduct(c echo.Context, productID int64) (*models.Product, error) {
	return u.repository.GetProduct(c, productID)
}

func (u usecase) GetProductWithTx(c echo.Context, productID int64) (*models.Product, error) {
	ac := c.(*util.CustomApplicationContext)

	tx := ac.MysqlSession.Begin()

	resp, err := u.repository.GetProduct(c, productID)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return resp, err
	}
	return resp, nil
}

// Get an addon from addon usecase
func (u usecase) GetAddonsByProductID(c echo.Context, productID int64) ([]models.Addon, error) {
	return u.addonUsecase.GetAllAddonsByProductID(c, productID)
}

// Public function
func CalculateDeliveryFee(basePrice int64) int64 {
	return basePrice * 20 / 100
}

// Private function
func calculateTax(basePrice int64) int64 {
	return basePrice / 10
}

// Take a product input and return with altered price
func AlterProductPrice(productInput models.Product) models.Product {
	productInput.Price += 1000
	return productInput
}
