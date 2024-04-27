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

	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		panic("SetTrustedProxies failed")
	}

	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.BookById)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/checkout", controllers.CheckoutBook)
	r.PATCH("/return", controllers.ReturnBook)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.GET("/classes", controllers.GetAll)
	r.POST("/classes", controllers.CreateClass)
	r.POST("/class_categories", controllers.CreateClassCategory)

	err := r.Run()
	if err != nil {
		return
	}

}
