package events

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/event")
	{
		public.GET("", controllers.GetEvent)
	}

	protected := route.Group("/event")
	protected.Use(middleware.RequireAuth)
	{
		protected.POST("", controllers.CreateEvent)
		protected.PATCH("", controllers.EditEvent)
		protected.DELETE("", controllers.DeleteEvent)
	}
}
