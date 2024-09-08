package models

type Todo struct {
	ID          int    `json:"id"`
	Task        string `json:"task"`
	IsCompleted bool   `json:"is_completed"`
}