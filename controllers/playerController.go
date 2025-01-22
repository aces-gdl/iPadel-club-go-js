package controllers

import (
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePlayer crea un nuevo jugador
func CreatePlayer(c *gin.Context) {
	var player models.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Create(&player)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
		return
	}

	c.JSON(http.StatusCreated, player)
}

// GetPlayer obtiene un jugador por su ID
func GetPlayer(c *gin.Context) {
	var player models.Player
	id := c.Param("id")

	if err := initializers.DB.Preload("Event").Preload("Person").Preload("Category").First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	c.JSON(http.StatusOK, player)
}

// GetAllPlayers obtiene todos los jugadores
func GetAllPlayers(c *gin.Context) {
	var players []models.Player
	initializers.DB.Preload("Event").Preload("Person").Preload("Category").Find(&players)
	c.JSON(http.StatusOK, players)
}

// UpdatePlayer actualiza un jugador existente
func UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player

	if err := initializers.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&player)
	c.JSON(http.StatusOK, player)
}

// DeletePlayer elimina un jugador
func DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player

	if err := initializers.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	initializers.DB.Delete(&player)
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}
