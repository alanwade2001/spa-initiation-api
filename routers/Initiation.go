package routers

import (
	"net/http"

	"github.com/alanwade2001/spa-initiation-api/generated/initiation"
	"github.com/alanwade2001/spa-initiation-api/types"
	"github.com/gin-gonic/gin"
)

// InitiationRouter s
type InitiationRouter struct {
	repositoryAPI types.RepositoryAPI
}

// NewInitiationRouter f
func NewInitiationRouter(repositoryAPI types.RepositoryAPI) types.InitiationAPI {

	initiationAPI := InitiationRouter{repositoryAPI}

	return &initiationAPI
}

// CreateInitiation f
func (cr *InitiationRouter) CreateInitiation(ctx *gin.Context) {
	init := new(initiation.InitiationModel)

	if err := ctx.BindJSON(init); err != nil {

		ctx.IndentedJSON(http.StatusUnprocessableEntity, err)

	} else if c1, err := cr.repositoryAPI.CreateInitiation(init); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusCreated, c1)
	}

}

// GetInitiation f
func (cr *InitiationRouter) GetInitiation(ctx *gin.Context) {
	initiationID := ctx.Param("id")
	if init, err := cr.repositoryAPI.GetInitiation(initiationID); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else if init == nil {
		ctx.Status(http.StatusNotFound)
	} else {
		ctx.IndentedJSON(http.StatusOK, init)
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
