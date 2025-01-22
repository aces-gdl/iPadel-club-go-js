package controllers

import (
	"fmt"
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {
	var body models.Persona

	result1 := c.Bind(&body)
	if result1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	if body.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Telefono y/o contraseña son requeridos...",
		})
		return
	}

	person := models.Persona{
		Model:       gorm.Model{},
		Phone:       body.Phone,
		Name:        body.Name,
		LastName:    body.LastName,
		BirthDate:   body.BirthDate,
		ClotingSize: body.ClotingSize,
	}

	if body.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Fallo al convertir password a hash...",
			})
			return
		}
		person.Password = string(hash)
	}

	result := initializers.DB.Create(&person)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var body models.Persona

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	if body.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Correo o clave invalido ...",
		})
		return
	}

	var person models.Persona
	initializers.DB.First(&person, "phone= ?", body.Phone)
	if person.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Correo o clave invalido ...",
		})
		return
	}

	if body.Password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "telefono o clave invalido ...",
			})
			return
		}

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": person.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear token ...",
		})
		return
	}

	strHost := strings.Split(c.Request.Host, ":")[0]

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", tokenString, 3600*8, "/", strHost, true, true)
	fmt.Println("Token generado: ", tokenString, " - ", strHost)
	c.JSON(http.StatusOK, gin.H{
		"message": "OK", "data": person,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout exitoso",
	})
}
