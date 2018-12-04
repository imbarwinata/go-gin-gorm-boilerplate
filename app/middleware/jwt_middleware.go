package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/app/models"
	"github.com/imbarwinata/go-gin-gorm-boilerplate/helpers/jwtauth"
)

func jwtAbort(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  401,
		"message": msg,
	})
	c.Abort()
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		models.Init()
		db := models.GetDB()

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			jwtAbort(c, "Authorization header dibutuhkan")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Boilerplate") {
			jwtAbort(c, "Authorization header tidak valid")
			return
		}

		claims, err := jwtauth.ParseToken(parts[1])
		if err != nil {
			jwtAbort(c, "Token tidak valid")
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			jwtAbort(c, "Tanggal kadaluarsa token telah berakhir")
			return
		}

		user := models.User{}

		if err := db.First(&user, claims.UserID).Error; err != nil {
			fmt.Println(err)
		}

		if user.ID != claims.UserID {
			jwtAbort(c, "Token tidak valid")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
