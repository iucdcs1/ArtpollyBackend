package models

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	Class    Class  `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;unique"`
	SUN      string `gorm:"default:null"` // JSON строка
	MON      string `gorm:"default:null"` // JSON строка
	TUE      string `gorm:"default:null"` // JSON строка
	WED      string `gorm:"default:null"` // JSON строка
	THU      string `gorm:"default:null"` // JSON строка
	FRI      string `gorm:"default:null"` // JSON строка
	SAT      string `gorm:"default:null"` // JSON строка
	Duration int    `gorm:"default:60"`
}
