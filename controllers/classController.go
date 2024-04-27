package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type classCategory struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type class struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Price       int           `json:"price"`
	CategoryID  uint          `json:"categoryId"`
	Category    classCategory `json:"category"`
}

func GetAll(ctx *gin.Context) {
	var categories []models.ClassCategory

	err := initializers.DB.Model(&models.ClassCategory{}).Preload("Classes").Find(&categories).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, categories)

}

func CreateClass(ctx *gin.Context) {
	var newClass class

	if err := ctx.BindJSON(&newClass); err != nil {
		return
	}

	newClass2 := models.Class{Title: newClass.Title, Description: newClass.Description, Price: newClass.Price, CategoryID: newClass.CategoryID}
	result := initializers.DB.Create(&newClass2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create class",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newClass2)
}

func CreateClassCategory(ctx *gin.Context) {
	var newClassCategory classCategory

	if err := ctx.BindJSON(&newClassCategory); err != nil {
		return
	}

	newClassCategory2 := models.ClassCategory{Title: newClassCategory.Title, Description: newClassCategory.Description}
	result := initializers.DB.Create(&newClassCategory2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create class category",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newClassCategory2)
}
