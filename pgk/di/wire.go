//+build wireinject

package di

import (
	"github.com/tonsV2/todo-go/pgk/configuration"
	"github.com/tonsV2/todo-go/pgk/storage/database"
	"github.com/tonsV2/todo-go/pgk/user"
)
import "github.com/google/wire"

type Environment struct {
	Configuration configuration.Configuration
	UserHandler   user.Handler
}

func ProvideEnvironment(
	configuration configuration.Configuration,
	userHandler user.Handler,
) Environment {
	return Environment{
		configuration,
		userHandler,
	}
}

func GetEnvironment() Environment {
	wire.Build(
		ProvideEnvironment,

		configuration.ProvideConfiguration,
		database.ProvideDatabase,

		user.ProvideRepository,
		user.ProvideService,
		user.ProvideHandler,
	)
	return Environment{}
}
