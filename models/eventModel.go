package models

import (
	"time"

	"gorm.io/gorm"
)

type EventType string

const (
	Tournament  EventType = "torneo"
	Flash       EventType = "relampago"
	League      EventType = "liga"
	KingOfCourt EventType = "rey de la cancha"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name" gorm:"not null"`
	ClubID      uint      `json:"club_id" gorm:"not null"`
	Club        Club      `json:"club" gorm:"foreignKey:ClubID"`
	StartDate   time.Time `json:"start_date" gorm:"not null"`
	EndDate     time.Time `json:"end_date" gorm:"not null"`
	EventType   EventType `json:"event_type" gorm:"not null"`
	ImageURL    string    `json:"image_url"`
	CourtCount  int       `json:"court_count" gorm:"not null"`
	Description string    `json:"description"`
}
