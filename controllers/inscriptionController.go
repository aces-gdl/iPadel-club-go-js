package controllers

import (
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateInscription crea una nueva inscripción
func CreateInscription(c *gin.Context) {
	type bodyT struct {
		EventID    uint `"json:event_id"`
		CategoryID uint `"json:category_id"`
		PersonID   uint `"json:person1_id"`
	}
	var body bodyT

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inscription models.Inscription
	inscription.InscriptionDate = time.Now()
	inscription.PersonID = body.PersonID
	inscription.EventID = body.EventID
	inscription.CategoryID = body.CategoryID

	result := initializers.DB.Create(&inscription)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inscription"})
		return
	}

	c.JSON(http.StatusCreated, inscription)
}

// GetInscription obtiene una inscripción por su ID
func GetInscription(c *gin.Context) {
	var inscription models.Inscription
	id := c.Param("id")

	if err := initializers.DB.Preload("Player").Preload("Event").Preload("Category").First(&inscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inscription not found"})
		return
	}

	c.JSON(http.StatusOK, inscription)
}

// GetAllInscriptions obtiene todas las inscripciones
func GetAllInscriptions(c *gin.Context) {
	var inscriptions []models.Inscription
	initializers.DB.Preload("Player").Preload("Event").Preload("Category").Find(&inscriptions)
	c.JSON(http.StatusOK, inscriptions)
}

// UpdateInscription actualiza una inscripción existente
func UpdateInscription(c *gin.Context) {
	id := c.Param("id")
	var inscription models.Inscription

	if err := initializers.DB.First(&inscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inscription not found"})
		return
	}

	if err := c.ShouldBindJSON(&inscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&inscription)
	c.JSON(http.StatusOK, inscription)
}

// DeleteInscription elimina una inscripción
func DeleteInscription(c *gin.Context) {
	id := c.Param("id")
	var inscription models.Inscription

	if err := initializers.DB.First(&inscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inscription not found"})
		return
	}

	initializers.DB.Delete(&inscription)
	c.JSON(http.StatusOK, gin.H{"message": "Inscription deleted successfully"})
}
