package middleware

import (
	"fmt"
	"nanam-yuk/auth"
	"nanam-yuk/initializers"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header)
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}

		var user auth.Auth
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}
}