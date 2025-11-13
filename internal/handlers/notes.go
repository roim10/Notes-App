package handlers

import (
	"notes-app/internal/models"
	"notes-app/internal/storage"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Получить все заметки
func GetNotes(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	notes := storage.GetAll()
	return c.JSON(notes)
}

// Создать заметку
func CreateNote(c *fiber.Ctx) error {
	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	created := storage.Add(note)

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.Status(201).JSON(created)
}

// Удалить заметку
func DeleteNote(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	storage.Delete(id)
	return c.SendStatus(204)
}
