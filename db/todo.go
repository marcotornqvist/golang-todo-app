package db

import (
	"database/sql"
	"time"

	"github.com/marcotornqvist/go-todo-app/models"
)

func (db Database) GetAllTodos() (*models.TodoList, error) {
	list := &models.TodoList{}

	rows, err := db.Conn.Query("SELECT * FROM todos ORDER BY ID DESC")

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)

		if err != nil {
			return list, err
		}

		list.Todos = append(list.Todos, todo)
	}

	return list, nil
}

func (db Database) AddTodo(todo *models.Todo) error {
	var id int
	var isCompleted bool
	var createdAt string
	var updatedAt string

	query := `INSERT INTO todos (title) VALUES ($1) RETURNING id, is_completed, created_at, updated_at`
	err := db.Conn.QueryRow(query, todo.Title).Scan(&id, &isCompleted, &createdAt, &updatedAt)

	if err != nil {
		return err
	}

	todo.ID = id
	todo.IsCompleted = isCompleted
	todo.UpdatedAt = updatedAt
	todo.CreatedAt = createdAt

	return nil
}

func (db Database) GetTodoById(todoId int) (models.Todo, error) {
	todo := models.Todo{}

	query := `SELECT * FROM todos WHERE id = $1;`
	row := db.Conn.QueryRow(query, todoId)

	switch err := row.Scan(&todo.ID, &todo.Title, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt); err {
	case sql.ErrNoRows:
		return todo, ErrNoMatch
	default:
		return todo, err
	}
}

func (db Database) UpdateTodo(todoId int, todoData models.Todo) (models.Todo, error) {
	todo := models.Todo{}

	query := `UPDATE todos SET title=$1, updated_at=$2 WHERE id=$3 RETURNING id, title, is_completed, created_at, updated_at;`
	err := db.Conn.QueryRow(query, todoData.Title, time.Now(), todoId).Scan(&todo.ID, &todo.Title, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return todo, ErrNoMatch
		}
		return todo, err
	}

	return todo, nil
}

func (db Database) DeleteTodo(todoId int) error {
	query := `DELETE FROM todos WHERE id = $1;`
	_, err := db.Conn.Exec(query, todoId)

	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) ToggleTodoCompleted(todoId int) (models.Todo, error) {
	todo := models.Todo{}

	query := `SELECT is_completed FROM todos WHERE id = $1;`
	err := db.Conn.QueryRow(query, todoId).Scan(&todo.IsCompleted)

	if err != nil {
		if err == sql.ErrNoRows {
			return todo, ErrNoMatch
		}

		return todo, err
	}

	

	query = `UPDATE todos SET is_completed=$1, updated_at=$2 WHERE id=$3 RETURNING id, title, created_at, updated_at;`
	err = db.Conn.QueryRow(query, !todo.IsCompleted, time.Now(), todoId).Scan(&todo.ID, &todo.Title, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return todo, ErrNoMatch
		}

		return todo, err
	}

	return todo, nil
}
