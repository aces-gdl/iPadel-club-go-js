package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	Cash     PaymentMethod = "efectivo"
	Transfer PaymentMethod = "transferencia"
	Card     PaymentMethod = "tarjeta"
)

type Inscription struct {
	gorm.Model
	PlayerID         uint          `json:"player_id" gorm:"not null"`
	Player           Player        `json:"player" gorm:"foreignKey:PlayerID"`
	EventID          uint          `json:"event_id" gorm:"not null"`
	Event            Event         `json:"event" gorm:"foreignKey:EventID"`
	CategoryID       uint          `json:"category_id" gorm:"not null"`
	Category         Category      `json:"category" gorm:"foreignKey:CategoryID"`
	PaymentMethod    PaymentMethod `json:"payment_method" gorm:"not null"`
	PaymentReference string        `json:"payment_reference"`
	InscriptionDate  time.Time     `json:"inscription_date" gorm:"not null"`
	PaymentConfirmed bool          `json:"payment_confirmed" gorm:"default:false"`
}
