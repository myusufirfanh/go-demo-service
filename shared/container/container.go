package container

import (
	"github.com/fgrosse/goldi"
	Config "github.com/myusufirfanh/go-demo-service/shared/config"
	Database "github.com/myusufirfanh/go-demo-service/shared/database"
)

// DefaultContainer returns default given depedency injections
func DefaultContainer() *goldi.Container {
	registry := goldi.NewTypeRegistry()

	config := make(map[string]interface{})
	container := goldi.NewContainer(registry, config)

	container.RegisterType("shared.config", Config.NewImmutableConfig)
	container.RegisterType("shared.mysql", Database.NewMysql, "@shared.config")
	return container
}
