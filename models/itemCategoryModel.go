package models

import "gorm.io/gorm"

type ItemCategory struct {
	gorm.Model
	Title       string
	Description string
}
