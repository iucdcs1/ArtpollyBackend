package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type schedule struct {
	ClassID  uint  `json:"id"`
	SUN      []int `json:"sun"`
	MON      []int `json:"mon"`
	TUE      []int `json:"tue"`
	WED      []int `json:"wed"`
	THU      []int `json:"thu"`
	FRI      []int `json:"fri"`
	SAT      []int `json:"sat"`
	Duration int   `json:"duration"`
}

func GetSchedule(ctx *gin.Context) {
	var schedules []models.Schedule

	err := initializers.DB.Model(&models.Schedule{}).Preload("Class").Find(&schedules).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, schedules)

}

func AddToSchedule(ctx *gin.Context) {
	var newSchedule schedule

	if err := ctx.BindJSON(&newSchedule); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var classToLink models.Class

	query := "ID = " + strconv.Itoa(int(newSchedule.ClassID))

	if err := initializers.DB.Model(&models.Class{}).Where(query).Find(&classToLink).Error; err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newScheduleModel := models.Schedule{
		Class:    classToLink,
		SUN:      newSchedule.SUN,
		MON:      newSchedule.MON,
		TUE:      newSchedule.TUE,
		WED:      newSchedule.WED,
		THU:      newSchedule.THU,
		FRI:      newSchedule.FRI,
		SAT:      newSchedule.SAT,
		Duration: newSchedule.Duration,
	}

	result := initializers.DB.Create(&newScheduleModel)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add schedule to mk " + newScheduleModel.Class.Title,
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newScheduleModel)
}

func UpdateSchedule(ctx *gin.Context) {
	var updatedSchedule schedule

	if err := ctx.BindJSON(&updatedSchedule); err != nil {
		return
	}

	var currentSchedule models.Schedule

	var classToLink models.Class

	query := "ID = " + strconv.Itoa(int(updatedSchedule.ClassID))

	if err := initializers.DB.Model(&models.Class{}).Where(query).Find(&classToLink).Error; err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := initializers.DB.Model(&models.Schedule{}).Where("Class = ?", classToLink).First(&currentSchedule).Error; err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedScheduleModel := models.Schedule{
		Class:    classToLink,
		SUN:      updatedSchedule.SUN,
		MON:      updatedSchedule.MON,
		TUE:      updatedSchedule.TUE,
		WED:      updatedSchedule.WED,
		THU:      updatedSchedule.THU,
		FRI:      updatedSchedule.FRI,
		SAT:      updatedSchedule.SAT,
		Duration: updatedSchedule.Duration,
	}

	updatedScheduleModel.ID = currentSchedule.ID

	result := initializers.DB.Save(&updatedScheduleModel)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedScheduleModel)
}

func DeleteSchedule(ctx *gin.Context) {
	var scheduleToRemove schedule

	if err := ctx.BindJSON(&scheduleToRemove); err != nil {
		return
	}

	var currentSchedule models.Schedule

	var classToLink models.Class

	query := "ID = " + strconv.Itoa(int(scheduleToRemove.ClassID))

	if err := initializers.DB.Model(&models.Class{}).Where(query).Find(&classToLink).Error; err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := initializers.DB.Model(&models.Schedule{}).Where("Class = ?", classToLink).First(&currentSchedule).Error; err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	entityToRemove := models.Schedule{}
	entityToRemove.ID = currentSchedule.ID

	result := initializers.DB.Unscoped().Delete(&entityToRemove)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, entityToRemove)
}
