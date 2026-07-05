package todo

import (
	"errors"

	"github.com/google/uuid"
	datasource "github.com/unim26/go_todo_server/internal/data_source"
	"github.com/unim26/go_todo_server/internal/models"
)

type todoRepository struct {
	source *datasource.DataSource
}

func NewtodoRepository(source *datasource.DataSource) *todoRepository {
	return &todoRepository{source: source}
}

// Create
func (tr *todoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	tr.source.Mu.Lock()
	defer tr.source.Mu.Unlock()

	if todo.Title == "" {
		return nil, errors.New("Title is required")
	}

	newId, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	todo.Id = newId
	todo.IsCompleted = false

	tr.source.TodoList = append(tr.source.TodoList, *todo)

	return todo, nil

}

// get all
func (tr *todoRepository) GetAll() ([]models.Todo, error) {
	tr.source.Mu.RLock()
	defer tr.source.Mu.RUnlock()

	return tr.source.TodoList, nil
}

// get by id
func (tr *todoRepository) GetById(id uuid.UUID) (*models.Todo, error) {
	if id.String() == "" {
		return nil, errors.New("not a valid id")
	}

	tr.source.Mu.RLock()
	defer tr.source.Mu.RUnlock()

	var matchedTodo *models.Todo
	for i := range tr.source.TodoList {
		if tr.source.TodoList[i].Id == id {
			matchedTodo = &tr.source.TodoList[i]
			break
		}
	}

	if matchedTodo == nil {
		return nil, errors.New("No todo found with id: " + id.String())
	}

	return matchedTodo, nil
}

// update
func (tr *todoRepository) Update(title string, id uuid.UUID) (*models.Todo, error) {
	if title == "" {
		return nil, errors.New("New title is required")
	}

	if id.String() == "" {
		return nil, errors.New("Not a valid id")
	}

	tr.source.Mu.Lock()
	defer tr.source.Mu.Unlock()

	var matchedTodo *models.Todo
	for i := range tr.source.TodoList {
		if tr.source.TodoList[i].Id == id {
			tr.source.TodoList[i].Title = title
			matchedTodo = &tr.source.TodoList[i]
			break
		}
	}

	if matchedTodo == nil {
		return nil, errors.New("No todo found with id: " + id.String())
	}

	return matchedTodo, nil
}

// toggle
func (tr *todoRepository) toggle(id uuid.UUID) error {
	if id.String() == "" {
		return errors.New("Not a valid id")
	}

	tr.source.Mu.Lock()
	defer tr.source.Mu.Unlock()

	found := false
	for i := range tr.source.TodoList {
		if tr.source.TodoList[i].Id == id {
			tr.source.TodoList[i].IsCompleted = !tr.source.TodoList[i].IsCompleted
			found = true
			break
		}
	}

	if !found {
		return errors.New("No todo found with id: " + id.String())
	}

	return nil

}

//delete
func (tr *todoRepository) Delete(id uuid.UUID) error {
	if id.String() == "" {
		return errors.New("Not a valid id")
	}

	tr.source.Mu.Lock()
	defer tr.source.Mu.Unlock()

	found := false
	for i := range tr.source.TodoList {
		if tr.source.TodoList[i].Id == id {
			tr.source.TodoList = append(tr.source.TodoList[:i], tr.source.TodoList[i + 1:]... )
			found = true
			break
		}
	}

	if !found {
		return errors.New("No todo found with id: " + id.String())
	}

	return nil
}
