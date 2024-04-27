package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
	Price       int
	CategoryID  uint `gorm:"default:null"`
}
