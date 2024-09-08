package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/ovelwe/todo-app-backend/config"
	"github.com/ovelwe/todo-app-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить .env файл")
	}

	config.Init()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173",
        AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
        AllowHeaders: "Content-Type, Authorization",
    }))
	routes.SetupRoutes(app)
	app.Listen(":3000")
}