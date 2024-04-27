package initializers

import (
	"artpollybackend/models"
)

func SyncDatabase() {
	if DB.AutoMigrate(&models.User{}) != nil {
		panic("Database models.User migration failed")
	}

	if DB.AutoMigrate(&models.Book{}) != nil {
		panic("Database models.Book migration failed")
	}

	if DB.AutoMigrate(&models.ClassCategory{}) != nil {
		panic("Database models.ClassCategory migration failed")
	}

	if DB.AutoMigrate(&models.Class{}) != nil {
		panic("Database models.Class migration failed")
	}

	if DB.AutoMigrate(&models.ItemCategory{}) != nil {
		panic("Database models.ItemCategory migration failed")
	}

	if DB.AutoMigrate(&models.Item{}) != nil {
		panic("Database models.Item migration failed")
	}

	if DB.AutoMigrate(&models.Event{}) != nil {
		panic("Database models.Event migration failed")
	}
}
