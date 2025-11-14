package service

import (
	"net/http"
	"time"
	"github.com/FudSy/DevVault/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (h *Handlers) CreateSnippet(c *gin.Context) {
	type Input struct {
	Title       string     `json:"title"`
	Code        string     `json:"code"`
	Language    models.Language   `json:"language"`
	Description string     `json:"description"`
	IsPublic    bool       `json:"is_public"`
	UserID      uuid.UUID `json:"-"`
	}

	uidValue, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non authorized"})
		return
	}

	input := Input{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while proccessing data"})
		log.Error().Msg("Error while binding JSON")
		return
	}

	snippet := models.Snippet{
		Title: input.Title,
		Code: input.Code,
		Language: input.Language,
		Description: input.Description,
		IsPublic: input.IsPublic,
		UserID: uidValue.(uuid.UUID),
		CreatedAt: time.Now(),
	}
	h.postgres.CreateSnippet(snippet)

	c.JSON(http.StatusOK, gin.H{"message": "snippet created"})
}