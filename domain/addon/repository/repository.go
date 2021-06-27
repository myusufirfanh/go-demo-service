package repository

import (
	"github.com/myusufirfanh/go-demo-service/domain/addon"
	"github.com/myusufirfanh/go-demo-service/models"

	"github.com/jinzhu/gorm"
)

type repoHandler struct {
	mysqlSess *gorm.DB
}

func NewAddonRepository(mysqlSess *gorm.DB) addon.Repository {
	return &repoHandler{
		mysqlSess: mysqlSess,
	}
}

func (r repoHandler) GetAddon(addonID int64) (*models.Addon, error) {
	res := &models.Addon{}
	db := r.mysqlSess.Where("ID = ?", addonID).First(&res)
	return res, db.Error
}

func (r repoHandler) GetAllAddonsByProductID(productID int64) ([]models.Addon, error) {
	res := []models.Addon{}
	db := r.mysqlSess.Where("product_id = ?", productID).Find(&res)
	return res, db.Error
}
