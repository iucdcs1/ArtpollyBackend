package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Title       string
	ImageURL    string
	StartDate   time.Time
	EndDate     time.Time
	Description string
}
