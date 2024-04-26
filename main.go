package main

import (
	"artpollybackend/controllers"
	"artpollybackend/initializers"
	"artpollybackend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.BookById)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/checkout", controllers.CheckoutBook)
	r.PATCH("/return", controllers.ReturnBook)

	err := r.Run("localhost:8080")
	if err != nil {
		return
	}

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	err = r.Run()
	if err != nil {
		return
	}

}
