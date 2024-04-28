package main

import (
	"artpollybackend/initializers"
	"artpollybackend/routes/classes"
	"artpollybackend/routes/events"
	"artpollybackend/routes/items"
	"artpollybackend/routes/users"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

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

	err := r.Run()
	if err != nil {
		return
	}

}
