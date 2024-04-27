package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title       string
	Description string
}
