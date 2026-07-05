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

func (tr *TodoRouter) RegisterTodoRoutes() {
	//create route
	http.HandleFunc("POST /todo", tr.handler.HandleCreate)
}
