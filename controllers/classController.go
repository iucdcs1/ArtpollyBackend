package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type classCategory struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type class struct {
	ID          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	ImageURL    string        `json:"image_url"`
	Price       int           `json:"price"`
	CategoryID  uint          `json:"categoryId"`
	Category    classCategory `json:"category"`
}

func GetCategorizedClasses(ctx *gin.Context) {
	var categories []models.ClassCategory

	err := initializers.DB.Model(&models.ClassCategory{}).Preload("Classes").Find(&categories).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, categories)

}

func GetUnassignedClasses(ctx *gin.Context) {
	var classes []models.Class

	err := initializers.DB.Model(&models.Class{}).Where("Category_ID IS NULL").Find(&classes).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, classes)

}

func CreateClass(ctx *gin.Context) {
	var newClass class

	if err := ctx.BindJSON(&newClass); err != nil {
		return
	}

	newClass2 := models.Class{
		Title:       newClass.Title,
		Description: newClass.Description,
		Price:       newClass.Price,
		ImageURL:    newClass.ImageURL,
		CategoryID:  newClass.CategoryID,
	}

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

	newClassCategory2 := models.ClassCategory{
		Title:       newClassCategory.Title,
		Description: newClassCategory.Description,
		ImageURL:    newClassCategory.ImageURL,
	}

	result := initializers.DB.Create(&newClassCategory2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create class category",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newClassCategory2)
}

func EditClass(ctx *gin.Context) {
	var newClass class

	if err := ctx.BindJSON(&newClass); err != nil {
		return
	}

	newClass2 := models.Class{
		Title:       newClass.Title,
		Description: newClass.Description,
		ImageURL:    newClass.ImageURL,
		Price:       newClass.Price,
		CategoryID:  newClass.CategoryID,
	}

	newClass2.ID = newClass.ID

	result := initializers.DB.Save(&newClass2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, newClass2)
}

func EditClassCategory(ctx *gin.Context) {
	var newCategory classCategory

	if err := ctx.BindJSON(&newCategory); err != nil {
		return
	}

	newCategory2 := models.ClassCategory{
		Title:       newCategory.Title,
		Description: newCategory.Description,
		ImageURL:    newCategory.ImageURL,
	}

	newCategory2.ID = newCategory.ID

	result := initializers.DB.Save(&newCategory2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, newCategory2)
}

func DeleteClass(ctx *gin.Context) {
	var classToRemove class

	if err := ctx.BindJSON(&classToRemove); err != nil {
		return
	}

	entityToRemove := models.Class{}
	entityToRemove.ID = classToRemove.ID

	result := initializers.DB.Unscoped().Delete(&entityToRemove)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, entityToRemove)
}

func DeleteClassCategory(ctx *gin.Context) {
	var categoryToRemove classCategory

	if err := ctx.BindJSON(&categoryToRemove); err != nil {
		return
	}

	entityToRemove := models.ClassCategory{}
	entityToRemove.ID = categoryToRemove.ID

	result := initializers.DB.Unscoped().Delete(&entityToRemove)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, entityToRemove)
}
