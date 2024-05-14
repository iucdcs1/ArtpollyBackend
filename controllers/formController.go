package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type formData struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", os.Getenv("TOKEN"))
}

func SendForm(ctx *gin.Context) {
	var data formData

	if err := ctx.BindJSON(&data); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	var response *http.Response

	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id": os.Getenv("CHAT_ID"),
		"text":    "Имя: " + data.Name + "\nТелефон: " + data.Phone + "\nПочта: " + data.Email,
	})

	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Telegram POST error": err.Error()})
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"TG Defer error:": err.Error()})
			return
		}
	}(response.Body)

	ctx.IndentedJSON(http.StatusOK, gin.H{"Telegram API response": response})
}
