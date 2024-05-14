package classes

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/class")
	{
		public.GET("", controllers.GetCategorizedClasses)
	}

	protected := route.Group("/class")
	protected.Use(middleware.RequireAuth)
	{
		protected.GET("/unassigned", controllers.GetUnassignedClasses)
		protected.POST("", controllers.CreateClass)
		protected.POST("/category", controllers.CreateClassCategory)
		protected.PATCH("", controllers.EditClass)
		protected.PATCH("/category", controllers.EditClassCategory)
		protected.DELETE("", controllers.DeleteClass)
		protected.DELETE("/category", controllers.DeleteClassCategory)
	}
}
