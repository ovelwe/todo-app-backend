package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ovelwe/todo-app-backend/config"
	"github.com/ovelwe/todo-app-backend/routes"
)

func main() {
	config.Init()

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}