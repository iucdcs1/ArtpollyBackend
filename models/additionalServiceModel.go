package models

import (
	"gorm.io/gorm"
)

type AdditionalService struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
	Price       int
	ImageURL    string
}
