package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ovelwe/todo-app-backend/controllers"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/todos", controllers.CreateTodo)
    app.Get("/todos", controllers.GetTodos)
    app.Delete("/todos/:id", controllers.DeleteTodo)
}