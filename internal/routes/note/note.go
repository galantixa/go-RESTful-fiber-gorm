package noteRoutes

import (
	noteHandler "github.com/galantixa/gofiber-gorm/internal/handlers/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	note.Post("/", noteHandler.CreateNotes)      // create note
	note.Get("/", noteHandler.GetNotes)          // get note
	note.Get("/:noteID", noteHandler.GetNote)    // get byID
	note.Put("/:noteID", noteHandler.UpdateNote) // update note byID
	note.Delete("/:noteID", noteHandler.DeleteNote)
}
