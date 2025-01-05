package main

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo
var idCounter = 1

func main() {
	app := fiber.New()

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		var newTodo Todo
		if err := c.BodyParser(&newTodo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
		newTodo.ID = idCounter
		idCounter++
		todos = append(todos, newTodo)
		return c.Status(fiber.StatusCreated).JSON(newTodo)
	})

	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		var updatedTodo Todo
		if err := c.BodyParser(&updatedTodo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
		for i, t := range todos {
			if t.ID == id {
				todos[i] = updatedTodo
				return c.JSON(updatedTodo)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON("Todo not found")
	})

	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.JSON("Todo deleted")
			}
		}
		return c.Status(fiber.StatusNotFound).JSON("Todo not found")
	})

	app.Listen(":8080")
}
