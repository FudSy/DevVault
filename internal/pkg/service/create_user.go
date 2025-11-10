package service

import (
	"net/http"
	"time"

	"github.com/FudSy/DevVault/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (h* Handlers) CreateUser(c *gin.Context) {
	type Input struct {
		Username  string `json:"username"`
		Password   string    `json:"password"`
		Email string `json:"email"`
	}
	input := Input{}
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while proccessing data"})
		log.Error().Msg("Error while binding JSON")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while proccessing data"})
		log.Error().Msg("Error while generating password hash")
		return
	}

	user := models.User {
		Username: input.Username,
		Password: string(hashedPassword),
		Email: input.Email,
		CreatedAt: time.Now(),
	}

	err = h.postgres.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while creating user"})
		log.Error().Msg("Error while creating user")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user has been registred"})
}