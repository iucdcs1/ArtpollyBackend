package forms

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/form")
	protected.Use(middleware.RequireAuth)
	{
		protected.POST("", controllers.SendForm)
	}
}
