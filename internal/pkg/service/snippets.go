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

func (h *Handlers) GetSnippet(c *gin.Context) {
	id := c.Query("id")
	
	type SnippetResponse struct {
	ID          uuid.UUID        `json:"id"`
	Title       string           `json:"title"`
	Code        string           `json:"code"`
	Language    models.Language  `json:"language"`
	Description string           `json:"description"`
	IsPublic    bool             `json:"is_public"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	}

	uidValue, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non authorized"})
		return
	}
	snippet, err := h.postgres.GetSnippet(uuid.MustParse(id), uidValue.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp := SnippetResponse{
	ID:          snippet.ID,
	Title:       snippet.Title,
	Code:        snippet.Code,
	Language:    snippet.Language,
	Description: snippet.Description,
	IsPublic:    snippet.IsPublic,
	CreatedAt:   snippet.CreatedAt,
	UpdatedAt: snippet.UpdatedAt,
}
	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) UpdateSnippet(c *gin.Context) {
	id := c.Query("id")
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
		UpdatedAt: time.Now(),
	}
	h.postgres.UpdateSnippet(snippet, uuid.MustParse(id), uidValue.(uuid.UUID))

	c.JSON(http.StatusOK, gin.H{"message": "snippet updated"})

}

func (h *Handlers) DeleteSnippet(c *gin.Context) {
	id := c.Query("id")

	uidValue, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non authorized"})
		return
	}
	h.postgres.DeleteSnippet(uuid.MustParse(id), uidValue.(uuid.UUID))

	c.JSON(http.StatusOK, gin.H{"message": "snippet deleted"})
}

func (h *Handlers) ListSnippet(c *gin.Context) {
	// TODO
}