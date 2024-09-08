package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func Init() {
	if godotenv.Load() != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	dbURL := os.Getenv("DATABASE_URL") 
	var err error
	
	connStr := dbURL
	DB, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
}