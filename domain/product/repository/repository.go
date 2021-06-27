package repository

import (
	"github.com/myusufirfanh/go-demo-service/domain/product"
	"github.com/myusufirfanh/go-demo-service/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type repoHandler struct {
	mysqlSess *gorm.DB
}

func NewProductRepository(mysqlSess *gorm.DB) product.Repository {
	return &repoHandler{
		mysqlSess: mysqlSess,
	}
}

func (r repoHandler) TxInsertProduct(c echo.Context, tx *gorm.DB, data *models.Product) (*models.Product, error) {
	db := tx.Create(&data)
	return data, db.Error
}

func (r repoHandler) GetProduct(c echo.Context, productID int64) (*models.Product, error) {
	res := &models.Product{}
	db := r.mysqlSess.Where("ID = ?", productID).First(&res)
	return res, db.Error
}
