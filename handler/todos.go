package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/marcotornqvist/go-todo-app/db"
	"github.com/marcotornqvist/go-todo-app/models"
)

type todoKey string

const todoIDKey todoKey = "todoID"

func todos(router chi.Router) {
	router.Get("/", getAllTodos)
	router.Post("/", createTodo)

	router.Route("/{todoId}", func(router chi.Router) {
		router.Use(TodoContext)
		router.Get("/", getTodo)
		router.Put("/", updateTodo)
		router.Delete("/", deleteTodo)
	})

	router.Route("/toggleTodo/{todoId}", func(router chi.Router) {
		router.Use(TodoContext)
		router.Put("/", toggleTodoCompleted)
	})
}

func TodoContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")

		if todoId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("todo ID is required")))
			return
		}

		id, err := strconv.Atoi(todoId)

		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid todo ID")))
		}

		ctx := context.WithValue(r.Context(), todoIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := dbInstance.GetAllTodos()

	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, todos); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}

	if err := render.Bind(r, todo); err != nil {

		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.AddTodo(todo); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, todo); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	todoID := r.Context().Value(todoIDKey).(int)
	fmt.Println(todoID)
	todo, err := dbInstance.GetTodoById(todoID)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &todo); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := r.Context().Value(todoIDKey).(int)
	todoData := models.Todo{}

	if err := render.Bind(r, &todoData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	todo, err := dbInstance.UpdateTodo(todoId, todoData)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &todo); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := r.Context().Value(todoIDKey).(int)
	err := dbInstance.DeleteTodo(todoId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func toggleTodoCompleted(w http.ResponseWriter, r *http.Request) {
	todoId := r.Context().Value(todoIDKey).(int)

	todo, err := dbInstance.ToggleTodoCompleted(todoId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &todo); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
