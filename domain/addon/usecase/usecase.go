package usecase

import (
	"github.com/labstack/echo"
	"github.com/myusufirfanh/go-demo-service/domain/addon"
	"github.com/myusufirfanh/go-demo-service/models"
)

type usecase struct {
	repository addon.Repository
}

func NewAddonUsecase(repository addon.Repository) addon.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u usecase) GetAddon(c echo.Context, addonID int64) (*models.Addon, error) {
	return u.repository.GetAddon(addonID)
}

func (u usecase) GetAllAddonsByProductID(c echo.Context, productID int64) ([]models.Addon, error) {
	return u.repository.GetAllAddonsByProductID(productID)
}
