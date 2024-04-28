package models

import "gorm.io/gorm"

type ItemCategory struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
	ImageURL    string
	Items       []Item `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
