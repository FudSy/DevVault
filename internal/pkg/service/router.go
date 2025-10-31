package service

import (
	"github.com/FudSy/DevVault/internal/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Router(postgres *postgres.DB) *gin.Engine {
	r := gin.Default()

	handlers := NewHandlers(postgres)

	r.POST("/register", handlers.CreateUser)

	return r
}