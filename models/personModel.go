package models

import (
	"time"

	"gorm.io/gorm"
)

type Persona struct {
	gorm.Model
	Phone       string    `json:"phone" gorm:"index:idx_phone,unique"`
	Password    string    `json:"password" `
	Name        string    `json:"name"`
	LastName    string    `json:"lastName"`
	BirthDate   time.Time `json:"birthDate"`
	ClotingSize string    `json:"clothingSize"`
}
