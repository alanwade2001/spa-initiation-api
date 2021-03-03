package main

import "github.com/gin-gonic/gin"

// RegisterService s
type RegisterService struct {
	Router        *gin.Engine
	initiationAPI InitiationAPI
}

// NewRegisterService f
func NewRegisterService(router *gin.Engine, initiationAPI InitiationAPI) RegisterAPI {

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
