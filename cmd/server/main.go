package main

import (
	"fmt"
	"net/http"

	"github.com/unim26/go_todo_server/cmd/router"
	datasource "github.com/unim26/go_todo_server/internal/data_source"
	"github.com/unim26/go_todo_server/internal/todo"
)

func main() {
	ds := datasource.NewDataSource()

	todoRepo := todo.NewtodoRepository(ds)
	todoHandler := todo.NewtodoHandler(todoRepo)
	todorouter := todo.NewtodoRouter(todoHandler)

	serverRouter := router.NewRouter(todorouter)

	serverMux := serverRouter.RegisterRouter()

	fmt.Println("Server is running on: https://api.unim.dev/api/v1")
	if err := http.ListenAndServe(":3000", serverMux); err != nil {
		fmt.Println("Error" + err.Error())
	}
}
