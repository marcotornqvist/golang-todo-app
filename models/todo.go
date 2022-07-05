package models

import (
	"fmt"
	"net/http"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
}

func (i *Todo) Bind(r *http.Request) error {
	if i.Title == "" {
		return fmt.Errorf("todo is a required field")
	}

	return nil
}

func (*TodoList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Todo) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
