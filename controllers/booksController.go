package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func GetBooks(ctx *gin.Context) {
	var books []models.Book
	initializers.DB.Find(&books)
	ctx.IndentedJSON(http.StatusOK, books)
}

func CreateBook(ctx *gin.Context) {
	var newBook book

	if err := ctx.BindJSON(&newBook); err != nil {
		return
	}

	newBook2 := models.Book{Title: newBook.Title, Author: newBook.Author, Quantity: newBook.Quantity}
	result := initializers.DB.Create(&newBook2)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create book",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newBook2)
}

func getBookByID(id string) (*models.Book, error) {
	var book models.Book
	initializers.DB.Where("id = ?", id).First(&book)

	if book.ID == 0 {
		return nil, errors.New("book not found")
	}

	return &book, nil
}

func BookById(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}

func CheckoutBook(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id"})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available..."})
		return
	}

	book.Quantity -= 1

	initializers.DB.Model(&book).Where("id = ?", book.ID).Update("quantity", book.Quantity).Commit()

	ctx.IndentedJSON(http.StatusOK, book)
}

func ReturnBook(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id"})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	book.Quantity += 1

	initializers.DB.Model(&book).Where("id = ?", book.ID).Update("quantity", book.Quantity).Commit()

	ctx.IndentedJSON(http.StatusOK, book)
}
