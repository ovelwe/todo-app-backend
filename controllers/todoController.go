package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ovelwe/todo-app-backend/models"
	"github.com/ovelwe/todo-app-backend/repo"
)

func CreateTodo(c *fiber.Ctx) error {
    var todo models.Todo
    if err := c.BodyParser(&todo); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Ошибка при создании задачи"})
    }
    err := repo.CreateTodo(todo)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Ошибка при добавлении задачи"})
    }
    return c.Status(201).JSON(todo)
}

func GetTodos(c *fiber.Ctx) error {
    todos, err := repo.GetTodos()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Ошибка при получении задач"})
    }
    return c.JSON(todos)
}

func DeleteTodo(c *fiber.Ctx) error {
    id, _ := c.ParamsInt("id")
    err := repo.DeleteTodo(id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Ошибка при удалении задачи"})
    }
    return c.SendStatus(204)
}