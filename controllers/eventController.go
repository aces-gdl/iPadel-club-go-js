package controllers

import (
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEvent crea un nuevo evento
func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Create(&event)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

// GetEvent obtiene un evento por su ID
func GetEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	if err := initializers.DB.Preload("Club").First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// GetAllEvents obtiene todos los eventos
func GetAllEvents(c *gin.Context) {
	var events []models.Event
	initializers.DB.Preload("Club").Find(&events)
	c.JSON(http.StatusOK, events)
}

// UpdateEvent actualiza un evento existente
func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event

	if err := initializers.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&event)
	c.JSON(http.StatusOK, event)
}

// DeleteEvent elimina un evento
func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event

	if err := initializers.DB.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	initializers.DB.Delete(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
