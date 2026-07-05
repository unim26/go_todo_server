package todo

import (
	"encoding/json"
	"net/http"

	"github.com/unim26/go_todo_server/internal/models"
	netservice "github.com/unim26/go_todo_server/internal/services/net_service"
)

type TodoHandler struct {
	repo Repository
}

func NewtodoHandler(repo Repository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

// handle create
func (th *TodoHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdTodo, err := th.repo.Create(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	netservice.SendResponse(w, http.StatusCreated, "Todo created successfully", createdTodo)

}

//handle getAll
func (th *TodoHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fetchedTodos, err := th.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	netservice.SendResponse(w, http.StatusOK, "Successfully fetched todos", fetchedTodos)
}

