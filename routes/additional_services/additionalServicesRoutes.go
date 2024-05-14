package additional_services

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/services")
	{
		public.GET("", controllers.GetServices)
	}

	protected := route.Group("/services")
	protected.Use(middleware.RequireAuth)
	{
		protected.POST("", controllers.CreateService)
		protected.PATCH("", controllers.UpdateService)
		protected.DELETE("", controllers.DeleteService)
	}
}
