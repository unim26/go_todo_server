package datasource

import (
	"sync"

	"github.com/unim26/go_todo_server/internal/models"
)

type DataSource struct {
	Mu       sync.RWMutex
	TodoList []models.Todo
}

func NewDataSource() *DataSource {
	return &DataSource{
		TodoList: make([]models.Todo, 0),
	}
}
