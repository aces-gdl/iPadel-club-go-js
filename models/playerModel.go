package models

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	EventID    uint     `json:"event_id" gorm:"not null;index:idx_player_event_category_persona,priority:1,unique"`
	Event      Event    `json:"event" gorm:"foreignKey:EventID"`
	PersonaID  uint     `json:"person_id" gorm:"not null;index:idx_player_event_category_persona,priority:3,unique"`
	Persona    Persona  `json:"persona" gorm:"foreignKey:PersonaID"`
	CategoryID uint     `json:"category_id" gorm:"not null;index:idx_player_event_category_persona,priority:2,unique"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}
