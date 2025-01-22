package controllers

import (
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory crea una nueva categoría
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategory obtiene una categoría por su ID
func GetCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := initializers.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetAllCategories obtiene todas las categorías
func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	initializers.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

// UpdateCategory actualiza una categoría existente
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := initializers.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&category)
	c.JSON(http.StatusOK, category)
}

// DeleteCategory elimina una categoría
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := initializers.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	initializers.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
