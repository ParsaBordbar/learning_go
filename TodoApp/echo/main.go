package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo
var idCounter = 1

func main() {
	e := echo.New()

	e.GET("/todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})

	e.POST("/todos", func(c echo.Context) error {
		var newTodo Todo
		if err := c.Bind(&newTodo); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		newTodo.ID = idCounter
		idCounter++
		todos = append(todos, newTodo)
		return c.JSON(http.StatusCreated, newTodo)
	})

	e.PUT("/todos/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var updatedTodo Todo
		if err := c.Bind(&updatedTodo); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		for i, t := range todos {
			if t.ID == id {
				todos[i] = updatedTodo
				return c.JSON(http.StatusOK, updatedTodo)
			}
		}
		return c.JSON(http.StatusNotFound, "Todo not found")
	})

	e.DELETE("/todos/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.JSON(http.StatusOK, "Todo deleted")
			}
		}
		return c.JSON(http.StatusNotFound, "Todo not found")
	})

	e.Start(":8080")
}
