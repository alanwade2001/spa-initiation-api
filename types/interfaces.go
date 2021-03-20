package types

import (
	"github.com/alanwade2001/spa-initiation-api/generated/initiation"
	"github.com/gin-gonic/gin"
)

// InitiationAPI i
type InitiationAPI interface {
	CreateInitiation(*gin.Context)
	GetInitiation(*gin.Context)
	GetInitiations(*gin.Context)
}

// ServerAPI i
type ServerAPI interface {
	Run() error
}

// RegisterAPI i
type RegisterAPI interface {
	Register() error
}

// RepositoryAPI i
type RepositoryAPI interface {
	CreateInitiation(c *initiation.InitiationModel) (*initiation.InitiationModel, error)
	GetInitiation(id string) (*initiation.InitiationModel, error)
	GetInitiations() ([]*initiation.InitiationModel, error)
}

// ConfigAPI si
type ConfigAPI interface {
	Load(path string) error
}
