package schedule

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/schedule")
	{
		public.GET("", controllers.GetSchedule)
	}

	protected := route.Group("/schedule")
	protected.Use(middleware.RequireAuth)
	{
		protected.POST("", controllers.AddToSchedule)
		protected.PATCH("", controllers.EditSchedule)
		protected.DELETE("", controllers.RemoveFromSchedule)
	}
}
