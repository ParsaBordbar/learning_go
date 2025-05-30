package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"go-fiber-todo/models"
	"go-fiber-todo/database"
)


func GetTodos(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT * FROM todos")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to query todos"})
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.AssignedTo, &t.State)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan todo"})
		}
		todos = append(todos, t)
	}

	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	for _, t := range models.Todos {
		if t.ID == id {
			return c.Status(fiber.StatusOK).JSON(t)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
}

func NewTodo(c *fiber.Ctx) error {
	return CreateTodo(c)
}

func CreateTodo(c *fiber.Ctx) error {
    todo := new(models.Todo)

    if err := c.BodyParser(todo); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    if todo.Title == "" {
        return c.Status(fiber.StatusBadRequest).SendString("Title cannot be empty")
    }

    if todo.AssignedTo == "" {
        return c.Status(fiber.StatusBadRequest).SendString("AssignedTo cannot be empty")
    }

    if todo.State == "" {
         return c.Status(fiber.StatusBadRequest).SendString("State cannot be empty")
    }

    result, err := database.DB.Exec("INSERT INTO todos (title, assigned_to, state) VALUES (?, ?, ?)", todo.Title, todo.AssignedTo, todo.State)
    if err != nil {
        log.Printf("Error inserting todo: %v", err)
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to create todo")
    }

    id, err := result.LastInsertId()
    if err != nil {
        log.Printf("Error getting last insert ID: %v", err)
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to create todo")
    }
    todo.ID = int(id)

    return c.Status(fiber.StatusCreated).JSON(todo)
}


func CreateTodoTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"title" TEXT NOT NULL,
		"assigned_to" TEXT NOT NULL,
		"state" TEXT NOT NULL
	);`
	statement, err := database.DB.Prepare(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	statement.Exec()
	log.Println("Table created successfully")
}
