package repo

import (
	"context"

	"github.com/ovelwe/todo-app-backend/config"
	"github.com/ovelwe/todo-app-backend/models"
)

func CreateTodo(todo models.Todo) error {
	_, err := config.DB.Exec(context.Background(), "INSERT INTO todos (task) VALUES ($1)", todo.Task)
	return err
}

func GetTodos() ([]models.Todo, error) {
    rows, err := config.DB.Query(context.Background(), "SELECT id, task FROM todos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var todo models.Todo
        if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

func DeleteTodo(id int) error {
	_, err := config.DB.Exec(context.Background(), "DELETE FROM todos WHERE id=$1", id)
	return err
}