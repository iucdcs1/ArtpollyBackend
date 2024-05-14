package users

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/")
	{
		public.POST("signup", controllers.Signup)
		public.POST("login", controllers.Login)
	}

	protected := route.Group("/")
	protected.Use(middleware.RequireAuth)
	{
		protected.GET("validate", controllers.Validate)
		protected.GET("user", controllers.GetInfo)
	}
}
