package initializers

import (
	"artpollybackend/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
