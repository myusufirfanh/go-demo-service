package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialect mysql for gorm
	"github.com/myusufirfanh/go-demo-service/shared/config"
	gormtracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/jinzhu/gorm"
)

var (
	once sync.Once
)

type (

	// MysqlInterface is an interface that represent mysql methods in package database
	MysqlInterface interface {
		OpenMysqlConn() (*gorm.DB, error)
	}

	// database is a struct to map given struct
	database struct {
		SharedConfig config.ImmutableConfigInterface
	}
)

func (d *database) OpenMysqlConn() (*gorm.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		d.SharedConfig.GetQoalaDatabaseUser(),
		d.SharedConfig.GetQoalaDatabasePassword(),
		d.SharedConfig.GetQoalaDatabaseHost(),
		d.SharedConfig.GetQoalaDatabasePort(),
		d.SharedConfig.GetQoalaDatabaseName(),
	)

	db, err := gormtracer.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(16)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.LogMode(true)
	return db, nil
}

// NewMysql is an factory that implement of mysql database configuration
func NewMysql(config config.ImmutableConfigInterface) MysqlInterface {
	if config == nil {
		panic("[CONFIG] immutable config is required")
	}

	return &database{config}
}
