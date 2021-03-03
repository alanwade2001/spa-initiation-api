package main

import "github.com/gin-gonic/gin"

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
	CreateInitiation(c *Initiation) (*Initiation, error)
	GetInitiation(id string) (*Initiation, error)
	GetInitiations() (*Initiations, error)
}

// ConfigAPI si
type ConfigAPI interface {
	Load() error
}
