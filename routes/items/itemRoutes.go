package items

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/item")
	{
		public.GET("", controllers.GetCategorizedItems)
	}

	protected := route.Group("/item")
	protected.Use(middleware.RequireAuth)
	{
		protected.GET("/unassigned", controllers.GetUnassignedItems)
		protected.POST("", controllers.CreateItem)
		protected.POST("/category", controllers.CreateItemCategory)
		protected.PATCH("", controllers.EditItem)
		protected.PATCH("/category", controllers.EditItemCategory)
		protected.DELETE("", controllers.DeleteItem)
		protected.DELETE("/category", controllers.DeleteItemCategory)
	}
}
