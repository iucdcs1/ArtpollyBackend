package main

import (
	"artpollybackend/initializers"
	"artpollybackend/routes/additional_services"
	"artpollybackend/routes/classes"
	"artpollybackend/routes/contacts"
	"artpollybackend/routes/events"
	"artpollybackend/routes/forms"
	"artpollybackend/routes/items"
	"artpollybackend/routes/schedules"
	"artpollybackend/routes/users"
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	// Forms
	forms.SetupRouter(r)
	// Schedules
	schedules.SetupRouter(r)
	// Contacts
	contacts.SetupRouter(r)

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Failed to start Service:", err)
		return
	}

	log.Info("Backend server is running!")

}
