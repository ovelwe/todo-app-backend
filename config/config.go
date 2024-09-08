package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Init() {
	var err error
	connStr := "postgres://postgres:123321@localhost:5432/todo-app"
	DB, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
}