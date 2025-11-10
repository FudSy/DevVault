package service

import (
	"net/http"
	"time"

	"github.com/FudSy/DevVault/internal/pkg/jwt"
	"github.com/FudSy/DevVault/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (h *Handlers) CreateSnippet(c *gin.Context) {
	jwtToken, err := c.Cookie("Token")
	if err != nil {
		log.Debug().Msg("Missing \"Token\" cookie")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non authorized"})
		return 
	}

	claims, err := jwt.ValidateToken(jwtToken)
	if err != nil {
		log.Error().Msg("Error while validating token")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	token, err := jwt.RefreshToken(jwtToken, time.Hour*3)
	if err != nil {
		log.Error().Msg("Error while refreshing token")
	}

	c.SetCookie("Token", token, int(time.Hour*3), "/", "", false, false)

	type Input struct {
	Title       string     `json:"title"`
	Code        string     `json:"code"`
	Language    models.Language   `json:"language"`
	Description string     `json:"description"`
	IsPublic    bool       `json:"is_public"`
	UserID      uuid.UUID `json:"-"`
	}
	input := Input{UserID: uuid.MustParse(claims.UserID)}
	err = c.ShouldBindJSON(&input)
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
		UserID: input.UserID,
		CreatedAt: time.Now(),
	}
	h.postgres.CreateSnippet(snippet)

	c.JSON(http.StatusOK, gin.H{"message": "snippet created"})
}