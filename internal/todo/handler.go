package todo


type TodoRepository struct {
	repo Repository
}

func NewTodoRepository(repo Repository) *TodoRepository {
	return &TodoRepository{repo: repo}
}


