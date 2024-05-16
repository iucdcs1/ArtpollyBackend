package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type scheduleObject struct {
	ID        uint      `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	DayOfWeek int       `json:"day_of_week"`
	Class     class     `json:"class,omitempty"`
}

func GetSchedule(ctx *gin.Context) {
	var schedule []models.ScheduleObject

	err := initializers.DB.Model(&models.ScheduleObject{}).Preload("Class").Find(&schedule).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, schedule)
}

func AddToSchedule(ctx *gin.Context) {
	var newScheduleObject scheduleObject
	if err := ctx.BindJSON(&newScheduleObject); err != nil {
		return
	}

	newScheduleObject2 := models.ScheduleObject{
		StartTime: newScheduleObject.StartTime,
		EndTime:   newScheduleObject.EndTime,
		DayOfWeek: newScheduleObject.DayOfWeek,
		ClassID:   newScheduleObject.Class.ID,
	}

	result := initializers.DB.Create(&newScheduleObject2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Failed to create schedule object": result.Error.Error()})
	}

	ctx.IndentedJSON(http.StatusCreated, newScheduleObject2)
}

func RemoveFromSchedule(ctx *gin.Context) {
	var scheduleObjectToRemove scheduleObject

	if err := ctx.BindJSON(&scheduleObjectToRemove); err != nil {
		return
	}

	entityToRemove := models.ScheduleObject{}
	entityToRemove.ID = scheduleObjectToRemove.ID

	result := initializers.DB.Unscoped().Delete(&entityToRemove)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error deleting schedule object": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, entityToRemove)
}

func EditSchedule(ctx *gin.Context) {
	var updatedSchedule []scheduleObject

	if err := ctx.BindJSON(&updatedSchedule); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error binding schedule JSON": err.Error()})
		return
	}

	tx := initializers.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, updatedScheduleObj := range updatedSchedule {
		updatedScheduleModel := models.ScheduleObject{

			StartTime: updatedScheduleObj.StartTime,
			EndTime:   updatedScheduleObj.EndTime,
			DayOfWeek: updatedScheduleObj.DayOfWeek,
			ClassID:   updatedScheduleObj.Class.ID,
		}
		updatedScheduleModel.ID = updatedScheduleObj.ID

		result := tx.Save(&updatedScheduleModel)
		if result.Error != nil {
			tx.Rollback()
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error updating schedule": result.Error.Error()})
			return
		}
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
}
