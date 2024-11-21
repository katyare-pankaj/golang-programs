// todo_api.go
package main  
import (  
    "fmt"
    "net/http"
    "strconv"

    "github.com/go-swagger/go-swagger/httpkit/middleware"
    httptransport "github.com/go-swagger/go-swagger/httpkit/transport"
)

// NewAPI creates a new Todo API handler
func NewAPI(spec *Spec) *operations.TodoAPI {
    return operations.NewTodoAPI(spec)
}

func (a *API) GetTodosHandler(params operations.GetTodosParams) middleware.Responder {
    todos := make(ToDos, 0, len(todosMap))
    for _, todo := range todosMap {
        todos = append(todos, todo)
    }
    return operations.NewGetTodosOK().WithPayload(todos)
}

func (a *API) GetTodoHandler(params operations.GetTodoParams) middleware.Responder {
    id, err := strconv.Atoi(params.ID)
