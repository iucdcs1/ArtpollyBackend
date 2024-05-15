package contacts

import (
	"artpollybackend/controllers"
	"artpollybackend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/contacts")
	{
		public.GET("", controllers.GetContacts)
	}

	protected := route.Group("/contacts")
	protected.Use(middleware.RequireAuth)
	{
		protected.PATCH("", controllers.UpdateContacts)
	}
}
