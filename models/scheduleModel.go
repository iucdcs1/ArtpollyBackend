package models

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	Class    Class `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;unique"`
	SUN      []int `gorm:"default:null"`
	MON      []int `gorm:"default:null"`
	TUE      []int `gorm:"default:null"`
	WED      []int `gorm:"default:null"`
	THU      []int `gorm:"default:null"`
	FRI      []int `gorm:"default:null"`
	SAT      []int `gorm:"default:null"`
	Duration int   `gorm:"default:60"`
}
