package controllers

import (
	"fmt"
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchPerson busca personas por cualquier campo
func SearchPerson(c *gin.Context) {
	var persons []models.Persona
	query := initializers.DB.Model(&models.Persona{})

	// Obtener todos los parámetros de búsqueda
	searchParams := c.Request.URL.Query()

	// Iterar sobre los parámetros y agregar condiciones de búsqueda
	for key, values := range searchParams {
		if len(values) > 0 {
			query = query.Where(key+" LIKE ?", "%"+values[0]+"%")
		}
	}

	// Ejecutar la consulta
	result := query.Find(&persons)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search persons"})
		return
	}

	if len(persons) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No persons found"})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func IsPersonAvailable(c *gin.Context) {
	PersonID := c.DefaultQuery("PersonID", "0")
	Phone := c.DefaultQuery("Phone", "0")
	EventID := c.DefaultQuery("EventID", "0")
	CategoryID := c.DefaultQuery("CategoryID", "0")

	if PersonID == "0" && Phone == "0" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No PersonID or Phone provided"})
		return
	}

	// phone over PersonID
	var person models.Persona
	if Phone != "0" {
		result := initializers.DB.Debug().First(&person, "phone=?", Phone)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error seach by phone, " + result.Error.Error()})
			return
		}

	}

	if PersonID != "0" {
		result := initializers.DB.Debug().First(&person, "id=?", PersonID)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error seach by PersonID"})
			return
		}
	}

	type teamExtended struct {
		EventID    uint `json:"event_id"`
		CategoryID uint `json:"category_id"`
		TeamID     uint `json:"team_id"`
		Player1ID  uint `json:"player1_id"`
		Person1ID  uint `json:"person1_id"`
		Player2ID  uint `json:"player2_id"`
		Person2ID  uint `json:"person2_id"`
	}
	sqlStatement := fmt.Sprintf("SELECT * FROM event_teams WHERE event_id = %s AND category_id = %s AND (persona1_id = %d OR persona2_id = %d)", EventID, CategoryID, person.ID, person.ID)
	var teamX teamExtended

	result := initializers.DB.Raw(sqlStatement).Scan(&teamX)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}
	if teamX.TeamID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Person already belongs to a team", "IsAvailable": false, "data": person})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person available", "IsAvailable": true, "data": person})

}

func GetLoggedInPerson(c *gin.Context) {

	person, _ := c.Get("user")
	c.JSON(http.StatusOK, person)
}
