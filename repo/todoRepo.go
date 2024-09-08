package repo

import (
	"context"

	"github.com/ovelwe/todo-app-backend/config"
	"github.com/ovelwe/todo-app-backend/models"
)

func CreateTodo(todo models.Todo) error {
	_, err := config.DB.Exec(context.Background(), "INSERT INTO todos (task, is_completed) VALUES ($1, $2)", todo.Task, todo.IsCompleted)
	return err
}

func GetTodos() ([]models.Todo, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, task, is_completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Task, &todo.IsCompleted); err != nil {
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