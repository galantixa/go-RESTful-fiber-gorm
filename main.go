package main

import (
	"github.com/galantixa/gofiber-gorm/database"
	"github.com/galantixa/gofiber-gorm/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// start fiber app
	app := fiber.New()
	database.ConnectDB()
	// setup routes
	router.SetupRoutes(app)

	// listen fiber app on port 5000
	app.Listen(":5000").Error()
}
