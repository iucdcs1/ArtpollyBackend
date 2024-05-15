package initializers

import (
	"artpollybackend/models"
)

func SyncDatabase() {
	if DB.AutoMigrate(&models.User{}) != nil {
		panic("Database models.User migration failed")
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

	if DB.AutoMigrate(&models.AdditionalService{}) != nil {
		panic("Database models.AdditionalService migration failed")
	}

	if DB.AutoMigrate(&models.Contacts{}) != nil {
		panic("Database models.Contacts migration failed")
	}
}
