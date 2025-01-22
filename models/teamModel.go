package models

import (
	"gorm.io/gorm"
)

type TeamStatus string

const (
	TeamPending   TeamStatus = "pendiente"
	TeamConfirmed TeamStatus = "confirmada"
	TeamCancelled TeamStatus = "cancelada"
)

type Team struct {
	gorm.Model
	EventID    uint       `json:"event_id" gorm:"not null"`
	Event      Event      `json:"event" gorm:"foreignKey:EventID"`
	CategoryID uint       `json:"category_id" gorm:"not null"`
	Category   Category   `json:"category" gorm:"foreignKey:CategoryID"`
	Player1ID  uint       `json:"player1_id" `
	Player2ID  uint       `json:"player2_id" `
	Status     TeamStatus `json:"status" gorm:"not null;default:'pendiente'"`
	TeamName   string     `json:"team_name"`
}
