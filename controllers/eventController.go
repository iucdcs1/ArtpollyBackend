package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type event struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetEvent(ctx *gin.Context) {
	var events []models.Event

	err := initializers.DB.Model(&models.Event{}).Preload("Events").Find(&events).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, events)

}

func CreateEvent(ctx *gin.Context) {
	var newEvent event

	if err := ctx.BindJSON(&newEvent); err != nil {
		return
	}

	newEvent2 := models.Event{Title: newEvent.Title, Description: newEvent.Description}
	result := initializers.DB.Create(&newEvent2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create event",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newEvent2)
}

func EditEvent(ctx *gin.Context) {
	var newEvent event

	if err := ctx.BindJSON(&newEvent); err != nil {
		return
	}

	newEvent2 := models.Event{Title: newEvent.Title, Description: newEvent.Description}
	newEvent2.ID = newEvent.ID

	result := initializers.DB.Save(&newEvent2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, newEvent2)
}

func DeleteEvent(ctx *gin.Context) {
	var eventToRemove event

	if err := ctx.BindJSON(&eventToRemove); err != nil {
		return
	}

	entityToRemove := models.Event{}
	entityToRemove.ID = eventToRemove.ID

	result := initializers.DB.Unscoped().Delete(&entityToRemove)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, entityToRemove)
}
