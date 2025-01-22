package controllers

import (
	"fmt"
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type myBody struct {
	EventID    uint
	CategoryID uint
	Person1ID  uint
	Player1ID  uint
	Person2ID  uint
	Player2ID  uint
	TeamName   string
	Status     models.TeamStatus `json:"Status" gorm:"default:'pendiente'"`
}

func CreateTeamWithTransaction(db *gorm.DB, body myBody) (message string, isError bool) {
	// Note the use of tx as the database handle once you are within a transaction

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return "Error creating transaction", true
	}

	var player models.Player
	// si no existe el jugador 1 entonces crearlo

	sqlStatement := fmt.Sprintf("SELECT * FROM players WHERE event_id = %d AND category_id = %d AND persona_id = %d", body.EventID, body.CategoryID, body.Person1ID)
	result := tx.Debug().Raw(sqlStatement).Scan(&player)
	if result.Error != nil {
		return "Error fetching player 1", true

	} else {
		if result.RowsAffected == 0 {
			player.EventID = body.EventID
			player.CategoryID = body.CategoryID
			player.PersonaID = body.Person1ID
			if err := tx.Create(&player).Error; err != nil {
				return "Failed to create player 1", true
			}
		}
	}

	body.Player1ID = player.ID
	var player2 models.Player
	// si el jugador 2 no esta vacio revisar que no existe y crearlo de ser necesario
	if body.Person2ID != 0 {
		sqlStatement = fmt.Sprintf("SELECT * FROM players WHERE event_id = %d AND category_id = %d AND persona_id = %d", body.EventID, body.CategoryID, body.Person2ID)
		result := tx.Debug().Raw(sqlStatement).Scan(&player2)
		if result.Error != nil {
			return "Error fetching player 2", true

		} else {
			if result.RowsAffected == 0 {
				player2.EventID = body.EventID
				player2.CategoryID = body.CategoryID
				player2.PersonaID = body.Person2ID
				if err := tx.Create(&player2).Error; err != nil {
					return "Failed to create player 2", true
				}
			}
		}
	}

	body.Player2ID = player2.ID

	type teamExtended struct {
		EventID    uint `json:"event_id"`
		CategoryID uint `json:"category_id"`
		TeamID     uint `json:"team_id"`
		Player1ID  uint `json:"player1_id"`
		Persona1ID uint `json:"person1_id"`
		Player2ID  uint `json:"player2_id"`
		Persona2ID uint `json:"person2_id"`
	}
	var teamX teamExtended
	sqlStatement = fmt.Sprintf("	select * from event_teams where event_id = %d and category_id = %d and (persona1_id in (%d,%d)  or persona2_id in (%d,%d))",
		body.EventID, body.CategoryID, body.Person1ID, body.Person2ID, body.Person1ID, body.Person2ID)

	result = tx.Debug().Raw(sqlStatement).Scan(&teamX)
	if result.Error != nil {
		return "Error fetching team", true
	} else {
		// revisar si es el mismo equipo
		if teamX.TeamID != 0 {
			// mismo equipo ya existe ->
			if (teamX.Player1ID == body.Player1ID && teamX.Player2ID == body.Player2ID) || teamX.Player1ID == body.Player2ID && teamX.Player2ID == body.Player1ID {
				return "Team with same players found", false
			}

			// player 1 ya existe en un equipo y P2 diferente
			if teamX.Player1ID == body.Player1ID && teamX.Player2ID != body.Player2ID {
				tx.Rollback()
				return "Team Player 1 already assigned", true
			}
			// player 2 ya existe en un equipo y P1 diferente
			if teamX.Player1ID == body.Player2ID && teamX.Player2ID != body.Player1ID {
				tx.Rollback()
				return "Team Player 2 already assigned", true
			}

		} else {
			// Equipo no Existe -> crearlo
			team := models.Team{
				EventID:    body.EventID,
				CategoryID: body.CategoryID,
				Player1ID:  body.Player1ID,
				Player2ID:  body.Player2ID,
				TeamName:   body.TeamName,
				Status:     body.Status,
			}
			if err := tx.Create(&team).Error; err != nil {
				tx.Rollback()
				return "Failed to create team", true
			}

		}
	}

	err := tx.Commit().Error.Error()
	if err != "" {
		return "err", true
	}
	return "Team creation success", false
}

// CreateTeam crea un nuevo equipo
func CreateTeam(c *gin.Context) {

	var body myBody

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Person1ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both player1_id and player2_id must be provided"})
		return
	}

	if body.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id must be provided"})
		return
	}
	if body.EventID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "event_id must be provided"})
		return
	}

	if err := initializers.DB.Debug().Where("id =?", body.CategoryID).Find(&models.Category{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}
	if err := initializers.DB.Debug().Where("id =?", body.EventID).Find(&models.Event{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event not found"})
		return
	}

	resultSTR, isError := CreateTeamWithTransaction(initializers.DB, body)

	if isError {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": resultSTR})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": resultSTR})
		return
	}

}

// GetTeam obtiene un equipo por su ID
func GetTeam(c *gin.Context) {
	var team models.Team
	id := c.Param("id")

	if err := initializers.DB.Preload("Event").Preload("Category").Preload("Person1").Preload("Person2").First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
}

// GetTeam obtiene un equipo por evento, categoría y persona
func GetTeamByPerson(c *gin.Context) {
	PersonID := c.DefaultQuery("PersonID", "0")
	EventID := c.DefaultQuery("EventID", "0")
	CategoryID := c.DefaultQuery("CategoryID", "0")

	type teamExtended struct {
		EventID    uint `json:"event_id"`
		CategoryID uint `json:"category_id"`
		TeamID     uint `json:"team_id"`
		Player1ID  uint `json:"player1_id"`
		Person1ID  uint `json:"person1_id"`
		Player2ID  uint `json:"player2_id"`
		Person2ID  uint `json:"person2_id"`
	}
	sqlStatement := fmt.Sprintf("SELECT * FROM event_teams WHERE event_id = %s AND category_id = %s AND (persona1_id = %s OR persona2 = %s)", EventID, CategoryID, PersonID, PersonID)
	var teamX teamExtended

	result := initializers.DB.Raw(sqlStatement).Scan(&teamX)
	if result.Error != nil && result.RowsAffected == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Person not found in teams"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": teamX, "IsAvailable": true})
}

// GetAllTeams obtiene todos los equipos
func GetAllTeams(c *gin.Context) {
	var teams []models.Team
	initializers.DB.Preload("Event").Preload("Category").Preload("Person1").Preload("Person2").Find(&teams)
	c.JSON(http.StatusOK, teams)
}

// UpdateTeam actualiza un equipo existente
func UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := initializers.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&team)
	c.JSON(http.StatusOK, team)
}

// DeleteTeam elimina un equipo
func DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := initializers.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	initializers.DB.Delete(&team)
	c.JSON(http.StatusOK, gin.H{"message": "Team deleted successfully"})
}
