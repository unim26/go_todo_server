package todo

import (
	"github.com/google/uuid"
	"github.com/unim26/go_todo_server/internal/models"
)

type Repository interface {
	//create todo
	Create(todo *models.Todo) (*models.Todo, error)

	//get all todo
	GetAll() ([]models.Todo, error)

	//get todo by id
	GetById(id uuid.UUID) (*models.Todo, error)

	//update todo
	Update(title string, id uuid.UUID) (*models.Todo, error)

	//toggle todo status
	toggle(id uuid.UUID) error

	//delete todo
	Delete(id uuid.UUID) error
}
