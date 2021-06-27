package main

import (
	"fmt"

	Config "github.com/myusufirfanh/go-demo-service/shared/config"

	Database "github.com/myusufirfanh/go-demo-service/shared/database"

	addonRepository "github.com/myusufirfanh/go-demo-service/domain/addon/repository"
	productRepository "github.com/myusufirfanh/go-demo-service/domain/product/repository"

	addonUsecase "github.com/myusufirfanh/go-demo-service/domain/addon/usecase"
	productUsecase "github.com/myusufirfanh/go-demo-service/domain/product/usecase"

	addonHandler "github.com/myusufirfanh/go-demo-service/domain/addon/delivery/http"
	productHandler "github.com/myusufirfanh/go-demo-service/domain/product/delivery/http"

	Container "github.com/myusufirfanh/go-demo-service/shared/container"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	container := Container.DefaultContainer()
	mysql := container.MustGet("shared.mysql").(Database.MysqlInterface)
	conf := container.MustGet("shared.config").(Config.ImmutableConfigInterface)

	mysqlSess, err := mysql.OpenMysqlConn()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open mysql connection: %s", err.Error())
		fmt.Println(msgError)
		panic(msgError)
	}

	productRepo := productRepository.NewProductRepository(mysqlSess)
	addonRepo := addonRepository.NewAddonRepository(mysqlSess)

	addonUcase := addonUsecase.NewAddonUsecase(addonRepo)
	productUcase := productUsecase.NewProductUsecase(productRepo, addonUcase)

	productHandler.AddProductHandler(e, productUcase)
	addonHandler.AddAddonHandler(e, addonUcase)

	e.Logger.Info(e.Start(fmt.Sprintf(":%d", conf.GetPort())))

}
