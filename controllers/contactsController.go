package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type contactsData struct {
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	WHours  string `json:"WHours"`
}

func GetContacts(ctx *gin.Context) {
	var contacts models.Contacts

	err := initializers.DB.Model(&models.Contacts{}).First(&contacts).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Contacts are empty": err.Error()})

		return
	}

	ctx.IndentedJSON(http.StatusOK, contacts)
}

func UpdateContacts(ctx *gin.Context) {
	var contacts contactsData

	if err := ctx.BindJSON(&contacts); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newContacts := models.Contacts{
		Email:   contacts.Email,
		Address: contacts.Address,
		Phone:   contacts.Phone,
		WHours:  contacts.WHours,
	}

	newContacts.ID = 1

	result := initializers.DB.Save(&newContacts)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, newContacts)
}
