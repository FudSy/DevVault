package service

import (
	"net/http"
	"time"

	"github.com/FudSy/DevVault/internal/pkg/jwt"
	"github.com/FudSy/DevVault/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (h* Handlers) Register(c *gin.Context) {
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

func (h* Handlers) Login(c *gin.Context) {
	type Input struct {
		Username  string `json:"username"`
		Password   string    `json:"password"`
	}
	input := Input{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while proccessing data"})
		log.Error().Msg("Error while binding JSON")
		return
	}

	userData, err := h.postgres.GetUserByUsername(input.Username)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad username or password"})
		log.Info().Msg("Record not found(bad username)")
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while proccessing data"})
		log.Error().Msg("Error while getting user by username")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad username or password"})
		log.Debug().Msg("Wrong user password")
		return
	} 

	jwtToken, _ := jwt.GenerateToken(userData.ID.String(), time.Hour*3)
	c.SetCookie("Token", jwtToken, int(time.Hour*3), "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "auth successful"})
}