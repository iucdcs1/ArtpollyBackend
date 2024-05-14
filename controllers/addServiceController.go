package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type additionalService struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       int    `json:"price"`
}

func GetServices(ctx *gin.Context) {
	var services []models.AdditionalService

	err := initializers.DB.Model(&models.AdditionalService{}).Find(&services).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, services)
}

func CreateService(ctx *gin.Context) {
	var newService additionalService

	if err := ctx.BindJSON(&newService); err != nil {
		return
	}

	newService2 := models.AdditionalService{
		Title:       newService.Title,
		Description: newService.Description,
		Price:       newService.Price,
		ImageURL:    newService.ImageURL,
	}
	result := initializers.DB.Create(&newService2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create class",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newService2)
}

func UpdateService(ctx *gin.Context) {
	var updatedService additionalService

	if err := ctx.BindJSON(&updatedService); err != nil {
		return
	}

	updatedServiceModel := models.AdditionalService{
		Title:       updatedService.Title,
		Description: updatedService.Description,
		ImageURL:    updatedService.ImageURL,
		Price:       updatedService.Price}

	updatedServiceModel.ID = updatedService.ID

	result := initializers.DB.Save(&updatedServiceModel)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedServiceModel)
}

func DeleteService(ctx *gin.Context) {
	var serviceToRemove additionalService

	if err := ctx.BindJSON(&serviceToRemove); err != nil {
		return
	}

	entityToRemove := models.AdditionalService{}
	entityToRemove.ID = serviceToRemove.ID

	result := initializers.DB.Unscoped().Delete(&entityToRemove)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, entityToRemove)
}
