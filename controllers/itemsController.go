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
}

type item struct {
	ID          uint         `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
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
