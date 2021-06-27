package config

import (
	"os"
	"sync"

	Error "github.com/myusufirfanh/go-demo-service/shared/error"
	"github.com/spf13/viper"
)

type (
	// ImmutableConfigInterface is an interface represent methods in config
	ImmutableConfigInterface interface {
		GetPort() int
		GetQoalaDatabaseHost() string
		GetQoalaDatabasePort() string
		GetQoalaDatabaseName() string
		GetQoalaDatabaseUser() string
		GetQoalaDatabasePassword() string
	}

	// im is a struct to mapping the structure of related value model
	im struct {
		Port                  int    `mapstructure:"PORT"`
		QoalaDatabaseHost     string `mapstructure:"QOALA_DATABASE_HOST"`
		QoalaDatabasePort     string `mapstructure:"QOALA_DATABASE_PORT"`
		QoalaDatabaseName     string `mapstructure:"QOALA_DATABASE_NAME"`
		QoalaDatabaseUser     string `mapstructure:"QOALA_DATABASE_USER"`
		QoalaDatabasePassword string `mapstructure:"QOALA_DATABASE_PASSWORD"`
	}
)

func (i *im) GetPort() int {
	return i.Port
}

func (i *im) GetQoalaDatabaseHost() string {
	return i.QoalaDatabaseHost
}

func (i *im) GetQoalaDatabasePort() string {
	return i.QoalaDatabasePort
}

func (i *im) GetQoalaDatabaseName() string {
	return i.QoalaDatabaseName
}

func (i *im) GetQoalaDatabaseUser() string {
	return i.QoalaDatabaseUser
}

func (i *im) GetQoalaDatabasePassword() string {
	return i.QoalaDatabasePassword
}

var (
	imOnce    sync.Once
	myEnv     map[string]string
	immutable im
)

// NewImmutableConfig is a factory that return of its config implementation
func NewImmutableConfig() ImmutableConfigInterface {
	imOnce.Do(func() {
		v := viper.New()
		appEnv, exists := os.LookupEnv("APP_ENV")
		if exists {
			if appEnv == "development" {
				v.SetConfigName("app.config.dev")
			} else if appEnv == "staging" {
				v.SetConfigName("app.config.staging")
			} else if appEnv == "production" {
				v.SetConfigName("app.config.prod")
			}
		} else {
			v.SetConfigName("app.config.local")
		}

		v.AddConfigPath(".")
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			Error.Wrap(500, "[SYS-001]", err, "[CONFIG][missing] Failed to read app.config.* file", "failed")
		}

		v.Unmarshal(&immutable)
	})

	return &immutable
}
