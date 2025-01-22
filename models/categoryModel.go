package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Description string `json:"description" gorm:"not null"`
	Level       string `json:"level" gorm:"not null"`
	Color       string `json:"color" gorm:"not null"`
	Primary     bool   `json:"primary"`
	ImageURL    string `json:"image_url"`
}
