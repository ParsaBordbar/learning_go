package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo
var idCounter = 1

func main() {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTodo.ID = idCounter
		idCounter++
		todos = append(todos, newTodo)
		c.JSON(http.StatusCreated, newTodo)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTodo Todo
		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, t := range todos {
			if id == string(t.ID) {
				todos[i] = updatedTodo
				c.JSON(http.StatusOK, updatedTodo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, t := range todos {
			if id == string(t.ID) {
				todos = append(todos[:i], todos[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	})

	r.Run(":8080")
}
