package main

import (
	"artpollybackend/initializers"
	"artpollybackend/routes/additional_services"
	"artpollybackend/routes/classes"
	"artpollybackend/routes/events"
	"artpollybackend/routes/items"
	"artpollybackend/routes/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		panic("SetTrustedProxies failed")
	}

	// Auth
	users.SetupRouter(r)
	// Classes
	classes.SetupRouter(r)
	// Items
	items.SetupRouter(r)
	// Events
	events.SetupRouter(r)
	// Services
	additional_services.SetupRouter(r)

	err := r.Run()
	if err != nil {
		return
	}

}
