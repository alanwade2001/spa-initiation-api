package routers

import (
	"github.com/alanwade2001/spa-initiation-api/types"
	"github.com/gin-gonic/gin"
)

// RegisterService s
type RegisterService struct {
	Router        *gin.Engine
	initiationAPI types.InitiationAPI
}

// NewRegisterService f
func NewRegisterService(router *gin.Engine, initiationAPI types.InitiationAPI) types.RegisterAPI {

	service := RegisterService{router, initiationAPI}
	return service

}

// Register f
func (rs RegisterService) Register() error {
	rs.Router.POST("/initiations", rs.initiationAPI.CreateInitiation)
	rs.Router.GET("/initiations", rs.initiationAPI.GetInitiations)
	rs.Router.GET("/initiations/:id", rs.initiationAPI.GetInitiation)

	return nil
}
