package service

import (
	"net/http"
	"time"

	"github.com/FudSy/DevVault/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h* Handlers) CreateUser(c *gin.Context) {
	type Input struct {
		Username  string `json:"username"`
		Password   string    `json:"password"`
		Email string `json:"email"`
	}
	input := Input{}
	c.ShouldBindJSON(&input)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User {
		Username: input.Username,
		Password: string(hashedPassword),
		Email: input.Email,
		CreatedAt: time.Now(),
	}
	h.postgres.CreateUser(&user)
	c.JSON(http.StatusOK, gin.H{"message": "user has been registred"})
}