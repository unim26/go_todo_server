package todo

import (
	"net/http"
)

type TodoRouter struct {
	handler *TodoHandler
}

func NewtodoRouter(handler *TodoHandler) *TodoRouter {
	return &TodoRouter{handler: handler}
}

func (tr *TodoRouter) RegisterTodoRoutes() *http.ServeMux {
	todoMux := http.NewServeMux()

	//create
	todoMux.HandleFunc("POST /", tr.handler.HandleCreate)

	//get all
	todoMux.HandleFunc("GET /", tr.handler.HandleGetAll)

	//get by id
	todoMux.HandleFunc("GET /{id}", tr.handler.GetById)

	return todoMux
}
