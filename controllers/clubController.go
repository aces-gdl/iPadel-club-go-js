package controllers

import (
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateClub crea un nuevo club
func CreateClub(c *gin.Context) {
	var club models.Club
	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Create(&club)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create club"})
		return
	}

	c.JSON(http.StatusCreated, club)
}

// GetClub obtiene un club por su ID
func GetClub(c *gin.Context) {
	var club models.Club
	id := c.Param("id")

	if err := initializers.DB.First(&club, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Club not found"})
		return
	}

	c.JSON(http.StatusOK, club)
}

// GetAllClubs obtiene todos los clubes
func GetAllClubs(c *gin.Context) {
	var clubs []models.Club
	initializers.DB.Find(&clubs)
	c.JSON(http.StatusOK, clubs)
}

// UpdateClub actualiza un club existente
func UpdateClub(c *gin.Context) {
	id := c.Param("id")
	var club models.Club

	if err := initializers.DB.First(&club, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Club not found"})
		return
	}

	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&club)
	c.JSON(http.StatusOK, club)
}

// DeleteClub elimina un club
func DeleteClub(c *gin.Context) {
	id := c.Param("id")
	var club models.Club

	if err := initializers.DB.First(&club, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Club not found"})
		return
	}

	initializers.DB.Delete(&club)
	c.JSON(http.StatusOK, gin.H{"message": "Club deleted successfully"})
}
