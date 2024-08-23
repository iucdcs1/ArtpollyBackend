package schedules

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/schedule")
	protected.Use(middleware.RequireAuth)
	{
		protected.POST("", controllers.AddToSchedule)
		protected.PATCH("", controllers.UpdateSchedule)
		protected.DELETE("", controllers.DeleteSchedule)
	}

	public := route.Group("/schedule")
	{
		public.GET("", controllers.GetSchedule)
	}
}
