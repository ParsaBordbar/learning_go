package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-fiber-todo/database"
	"go-fiber-todo/routes"
)


func main() {
	app := fiber.New()
	app.Use(logger.New())
	database.Connect()
	routes.SetupRoutes(app)
	 
	log.Fatal(app.Listen(":3000"))
	defer database.DB.Close()
}
	