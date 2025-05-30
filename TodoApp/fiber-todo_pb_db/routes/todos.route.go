package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-todo/handlers"
)

func SetupRoutes(app *fiber.App) {
	todoGroup := app.Group("/todos")
	todoGroup.Get("/", handlers.GetTodos)
	todoGroup.Get("/:id", handlers.GetTodo)
	todoGroup.Post("/new-todo", handlers.NewTodo)
}