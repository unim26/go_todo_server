package todo

import (
	"errors"

	"github.com/google/uuid"
	datasource "github.com/unim26/go_todo_server/internal/data_source"
	"github.com/unim26/go_todo_server/internal/models"
)

type todoDataSource struct {
	source *datasource.DataSource
}

func NewTodoDataSource(source *datasource.DataSource) *todoDataSource {
	return &todoDataSource{source: source}
}

// Create
func (ds *todoDataSource) Create(todo models.Todo) (*models.Todo, error) {
	ds.source.Mu.Lock()
	defer ds.source.Mu.Unlock()

	if todo.Title == "" {
		return nil, errors.New("Title is required")
	}

	newId, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	todo.Id = newId
	todo.IsCompleted = false

	ds.source.TodoList = append(ds.source.TodoList, todo)

	return &todo, nil

}

// get all
func (ds *todoDataSource) GetAll() ([]models.Todo, error) {
	ds.source.Mu.RLock()
	defer ds.source.Mu.RUnlock()

	return ds.source.TodoList, nil
}

// get by id
func (ds *todoDataSource) GetById(id uuid.UUID) (*models.Todo, error) {
	if id.String() == "" {
		return nil, errors.New("not a valid id")
	}

	ds.source.Mu.RLock()
	defer ds.source.Mu.RUnlock()

	var matchedTodo *models.Todo
	for i := range ds.source.TodoList {
		if ds.source.TodoList[i].Id == id {
			matchedTodo = &ds.source.TodoList[i]
			break
		}
	}

	if matchedTodo == nil {
		return nil, errors.New("No todo found with id: " + id.String())
	}

	return matchedTodo, nil
}

// update
func (ds *todoDataSource) Update(title string, id uuid.UUID) (*models.Todo, error) {
	if title == "" {
		return nil, errors.New("New title is required")
	}

	if id.String() == "" {
		return nil, errors.New("Not a valid id")
	}

	ds.source.Mu.Lock()
	defer ds.source.Mu.Unlock()

	var matchedTodo *models.Todo
	for i := range ds.source.TodoList {
		if ds.source.TodoList[i].Id == id {
			ds.source.TodoList[i].Title = title
			matchedTodo = &ds.source.TodoList[i]
			break
		}
	}

	if matchedTodo == nil {
		return nil, errors.New("No todo found with id: " + id.String())
	}

	return matchedTodo, nil
}

// toggle
func (ds *todoDataSource) toggle(id uuid.UUID) error {
	if id.String() == "" {
		return errors.New("Not a valid id")
	}

	ds.source.Mu.Lock()
	defer ds.source.Mu.Unlock()

	found := false
	for i := range ds.source.TodoList {
		if ds.source.TodoList[i].Id == id {
			ds.source.TodoList[i].IsCompleted = !ds.source.TodoList[i].IsCompleted
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
func (ds *todoDataSource) Delete(id uuid.UUID) error {
	if id.String() == "" {
		return errors.New("Not a valid id")
	}

	ds.source.Mu.Lock()
	defer ds.source.Mu.Unlock()

	found := false
	for i := range ds.source.TodoList {
		if ds.source.TodoList[i].Id == id {
			ds.source.TodoList = append(ds.source.TodoList[:i], ds.source.TodoList[i + 1:]... )
			found = true
			break
		}
	}

	if !found {
		return errors.New("No todo found with id: " + id.String())
	}

	return nil
}
