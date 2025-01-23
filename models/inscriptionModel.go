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
	PersonID         uint          `json:"person_id" gorm:"not null;index:idx_inscription_person_event_category,priority:3,unique"`
	Person           Persona       `json:"person" gorm:"foreignKey:PersonID"`
	EventID          uint          `json:"event_id" gorm:"not null;index:idx_inscription_person_event_category,priority:1,unique"`
	Event            Event         `json:"event" gorm:"foreignKey:EventID"`
	CategoryID       uint          `json:"category_id" gorm:"not null;index:idx_inscription_person_event_category,priority:2,unique"`
	Category         Category      `json:"category" gorm:"foreignKey:CategoryID"`
	PaymentMethod    PaymentMethod `json:"payment_method" gorm:"not null"`
	PaymentReference string        `json:"payment_reference"`
	InscriptionDate  time.Time     `json:"inscription_date" gorm:"not null"`
	PaymentConfirmed bool          `json:"payment_confirmed" gorm:"default:false"`
}
