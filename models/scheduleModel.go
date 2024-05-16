package models

import (
	"gorm.io/gorm"
	"time"
)

type ScheduleObject struct {
	gorm.Model
	StartTime time.Time
	EndTime   time.Time
	DayOfWeek int
	ClassID   uint
	Class     Class `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
