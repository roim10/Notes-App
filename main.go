package main

import (
	"log"

	"notes-app/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Включаем логирование всех запросов
	app.Use(logger.New())

	// Включаем CORS для всех источников
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // разрешаем любые источники
		AllowMethods: "GET,POST,DELETE,OPTIONS",
	}))

	// Маршруты API
	app.Get("/notes", handlers.GetNotes)          // получить все заметки
	app.Post("/notes", handlers.CreateNote)       // создать заметку
	app.Delete("/notes/:id", handlers.DeleteNote) // удалить заметку по id

	// Минимальный маршрут для проверки сервера
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Сервер работает!")
	})

	// Запуск сервера на порту 8080
	log.Fatal(app.Listen(":8080"))
}
