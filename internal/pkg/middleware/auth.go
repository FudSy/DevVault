package middleware

import (
	"net/http"
	"time"

	"github.com/FudSy/DevVault/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "non authorized"})
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}

		newToken, err := jwt.RefreshToken(token, time.Hour*3)
		if err == nil {
			c.SetCookie("Token", newToken, int(time.Hour*3), "/", "", false, false)
		}

		uid := uuid.MustParse(claims.UserID)
		c.Set("user_id", uid)

		c.Next()
	}
}
