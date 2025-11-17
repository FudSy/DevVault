package service

import (
	"github.com/FudSy/DevVault/internal/pkg/middleware"
	"github.com/FudSy/DevVault/internal/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Router(postgres *postgres.DB) *gin.Engine {
	r := gin.Default()

	handlers := NewHandlers(postgres)

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/", middleware.Auth())
	{
		auth.POST("/snippet", handlers.CreateSnippet)
		auth.GET("/snippet", handlers.GetSnippet)
		auth.PUT("/snippet", handlers.UpdateSnippet)
		auth.DELETE("/snippet", handlers.DeleteSnippet)
	}

	return r
}