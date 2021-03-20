//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/alanwade2001/spa-initiation-api/repositories"
	"github.com/alanwade2001/spa-initiation-api/routers"
	"github.com/alanwade2001/spa-initiation-api/services"
	"github.com/alanwade2001/spa-initiation-api/types"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitialiseServerAPI() types.ServerAPI {
	wire.Build(
		gin.Default,
		repositories.NewMongoRepository,
		routers.NewInitiationRouter,
		routers.NewRegisterService,
		services.NewConfigService,
		NewServer,
	)

	return &Server{}
}
