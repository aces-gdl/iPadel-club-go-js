package middleware

import (
	"fmt"
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"
	"strings"

	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	if os.Getenv("ENABLED_SECURITY") == "NO" {
		c.Next()
		return
	}

	tokenString, err := c.Cookie("jwt")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var person models.Persona
		initializers.DB.First(&person, "id = ?", claims["sub"])
		if person.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", person)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
