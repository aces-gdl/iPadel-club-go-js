package models

import (
	"gorm.io/gorm"
)

type Club struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Website     string `json:"website"`
	Description string `json:"description"`
	TotalCourts int    `json:"total_courts"`
	ImageURL    string `json:"image_url"`
}
