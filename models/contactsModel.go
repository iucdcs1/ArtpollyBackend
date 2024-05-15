package models

import (
	"gorm.io/gorm"
)

type Contacts struct {
	gorm.Model
	Email   string
	Address string
	Phone   string
	WHours  string
}
