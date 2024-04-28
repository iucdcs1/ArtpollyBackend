package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type itemCategory struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type item struct {
	ID          uint         `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	ImageURL    string       `json:"image_url"`
	Price       int          `json:"price"`
	CategoryID  uint         `json:"categoryId"`
	Category    itemCategory `json:"category"`
}

func GetCategorizedItems(ctx *gin.Context) {
	var categories []models.ItemCategory

	err := initializers.DB.Model(&models.ItemCategory{}).Preload("Items").Find(&categories).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, categories)
}

func GetUnassignedItems(ctx *gin.Context) {
	var items []models.Item

	err := initializers.DB.Model(&models.Item{}).Where("Category_ID IS NULL").Find(&items).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, items)
}

func CreateItem(ctx *gin.Context) {
	var newItem item

	if err := ctx.BindJSON(&newItem); err != nil {
		return
	}

	newItem2 := models.Item{Title: newItem.Title, Description: newItem.Description, ImageURL: newItem.ImageURL, Price: newItem.Price, CategoryID: newItem.CategoryID}
	result := initializers.DB.Create(&newItem2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create class",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newItem2)
}

func CreateItemCategory(ctx *gin.Context) {
	var newItemCategory itemCategory

	if err := ctx.BindJSON(&newItemCategory); err != nil {
		return
	}

	newItemCategory2 := models.ItemCategory{Title: newItemCategory.Title, Description: newItemCategory.Description, ImageURL: newItemCategory.ImageURL}
	result := initializers.DB.Create(&newItemCategory2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create class category",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newItemCategory2)
}

func EditItem(ctx *gin.Context) {
	var newItem item

	if err := ctx.BindJSON(&newItem); err != nil {
		return
	}

	newItem2 := models.Item{Title: newItem.Title, Description: newItem.Description, ImageURL: newItem.ImageURL, Price: newItem.Price, CategoryID: newItem.CategoryID}
	newItem2.ID = newItem.ID

	result := initializers.DB.Save(&newItem2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, newItem2)
}

func EditItemCategory(ctx *gin.Context) {
	var newCategory itemCategory

	if err := ctx.BindJSON(&newCategory); err != nil {
		return
	}

	newCategory2 := models.ItemCategory{Title: newCategory.Title, Description: newCategory.Description, ImageURL: newCategory.ImageURL}
	newCategory2.ID = newCategory.ID

	result := initializers.DB.Save(&newCategory2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, newCategory2)
}

func DeleteItem(ctx *gin.Context) {
	var itemToRemove class

	if err := ctx.BindJSON(&itemToRemove); err != nil {
		return
	}

	entityToRemove := models.Item{}
	entityToRemove.ID = itemToRemove.ID

	result := initializers.DB.Unscoped().Delete(&entityToRemove)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, entityToRemove)
}

func DeleteItemCategory(ctx *gin.Context) {
	var categoryToRemove itemCategory

	if err := ctx.BindJSON(&categoryToRemove); err != nil {
		return
	}

	entityToRemove := models.ItemCategory{}
	entityToRemove.ID = categoryToRemove.ID

	result := initializers.DB.Unscoped().Delete(&entityToRemove)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, entityToRemove)
}
