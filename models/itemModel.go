package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
	Price       int
	ItemID      uint
	Category    ItemCategory `gorm:"foreignKey:ItemID"`
}
