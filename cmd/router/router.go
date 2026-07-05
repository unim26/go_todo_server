package router

import (
	"encoding/json"
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

func (r *router) RegisterRouter() *http.ServeMux {
	basePrefix := "/api/v1"
	routerMux := http.NewServeMux()

	routerMux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status": "ok", "message": "Server is healthy",
		})
	})

	todoMux := r.tr.RegisterTodoRoutes()

	routerMux.Handle(basePrefix+"/todo/", http.StripPrefix(basePrefix+"/todo", todoMux))

	return routerMux
}
