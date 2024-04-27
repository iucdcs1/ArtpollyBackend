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

	// Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	// Classes
	r.GET("/classes", controllers.GetCategorizedClasses)
	r.GET("/unassigned_classes", controllers.GetUnassignedClasses)
	r.POST("/classes", controllers.CreateClass)
	r.POST("/class_categories", controllers.CreateClassCategory)
	r.PATCH("/class_edit", controllers.EditClass)
	r.PATCH("/class_category_edit", controllers.EditClassCategory)
	r.DELETE("/delete_class", controllers.DeleteClass)
	r.DELETE("/delete_class_category", controllers.DeleteClassCategory)
	// Items
	r.GET("/items", controllers.GetCategorizedItems)
	r.GET("/unassigned_items", controllers.GetUnassignedItems)
	r.POST("/items", controllers.CreateItem)
	r.POST("/item_categories", controllers.CreateItemCategory)
	r.PATCH("/item_edit", controllers.EditItem)
	r.PATCH("/item_category_edit", controllers.EditItemCategory)
	r.DELETE("/delete_item", controllers.DeleteItem)
	r.DELETE("/delete_item_category", controllers.DeleteItemCategory)

	err := r.Run()
	if err != nil {
		return
	}

}
