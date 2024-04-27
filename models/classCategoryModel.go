package models

import "gorm.io/gorm"

type ClassCategory struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
	Classes     []Class `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
