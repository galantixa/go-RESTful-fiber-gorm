package router

import (
	noteRoutes "github.com/galantixa/gofiber-gorm/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	noteRoutes.SetupNoteRoutes(api)
}
