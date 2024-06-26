package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
	Price       int
	ImageURL    string
	CategoryID  uint `gorm:"default:null"`
}
