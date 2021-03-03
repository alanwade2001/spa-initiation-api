package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitiationRouter s
type InitiationRouter struct {
	repositoryAPI RepositoryAPI
}

// NewInitiationRouter f
func NewInitiationRouter(repositoryAPI RepositoryAPI) InitiationAPI {

	initiationAPI := InitiationRouter{repositoryAPI}

	return &initiationAPI
}

// CreateInitiation f
func (cr *InitiationRouter) CreateInitiation(ctx *gin.Context) {
	initiation := new(Initiation)

	if err := ctx.BindJSON(initiation); err != nil {

		ctx.IndentedJSON(http.StatusUnprocessableEntity, err)

	} else if c1, err := cr.repositoryAPI.CreateInitiation(initiation); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusCreated, c1)
	}

}

// GetInitiation f
func (cr *InitiationRouter) GetInitiation(ctx *gin.Context) {
	initiationID := ctx.Param("id")
	if initiation, err := cr.repositoryAPI.GetInitiation(initiationID); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, initiation)
	}
}

// GetInitiations f
func (cr *InitiationRouter) GetInitiations(ctx *gin.Context) {
	if initiations, err := cr.repositoryAPI.GetInitiations(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, initiations)
	}
}
