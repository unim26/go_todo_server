package router

import (
	"net/http"

	"github.com/unim26/go_todo_server/internal/todo"
)

type router struct {
	tr todo.TodoRouter
}

func NewRouter(todoRouter *todo.TodoRouter) *router {
	return &router{
		tr: *todoRouter,
	}
}

func (r *router) RegisterRouter() {
	http.HandleFunc("POST /note", r.tr.RegisterTodoRoutes())
}
