package initializers

import (
	"ipadel-club/models"
)

func SyncDatabase() {
	DB.AutoMigrate(
		&models.Persona{},
		&models.Club{},
		&models.Event{},
		&models.Category{},
		&models.Player{},
		&models.Inscription{},
		&models.Team{})
}
