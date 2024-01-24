package noteHandlers

import (
	"github.com/galantixa/gofiber-gorm/database"
	"github.com/galantixa/gofiber-gorm/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	db.Find(&notes)
	if len(notes) == 0 {
		return c.Status(404).JSONP(fiber.Map{"Status": "error", "message": "No notes present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// add uuid to the notes
	note.ID = uuid.New()
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	return c.JSONP(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

// get note handler

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	//read param noteID
	id := c.ParamsParser("noteID")
	//find note with ID
	db.Find(&note, "id = ?", id)

	// condition when error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}
	// return the note with ID
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

// update note handler

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}
	db := database.DB
	var note model.Note

	// read param noteID
	id := c.Params("noteID")
	// find the note by ID
	db.Find(&note, "id = ?", id)

	// return error if there is not found
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// store updated data
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Edit the note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	db.Save(&note)
	// return change updated data
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// read param noteid
	id := c.Params("noteID")

	// find notes with the ID
	db.Find(&note, "id == ?", id)

	// return error if id not found
	if note.ID != uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// delete note and return error if failed
	err := db.Delete(&note, "id == ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	// return success
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
