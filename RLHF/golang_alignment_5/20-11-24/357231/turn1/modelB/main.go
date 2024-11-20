// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	httptransport "github.com/go-swagger/go-swagger/httpkit/transport"
	"github.com/gorilla/mux"
)

// Todo struct defines the structure of a to-do item
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// ToDos array of Todo struct
type ToDos []Todo

// todosMap is a simple in-memory map to store to-do items (for demonstration purposes)
var todosMap map[int]Todo

func init() {
	todosMap = make(map[int]Todo)
	// Add some sample to-do items for demonstration
	todosMap[1] = Todo{ID: 1, Title: "Buy milk", Done: false}
	todosMap[2] = Todo{ID: 2, Title: "Learn Go", Done: true}
}

func main() {
	r := mux.NewRouter()
	api := NewAPI(swaggerSpec())

	// Register the API routes
	r.Handle("/", httptransport.New(api, api.GetTodosHandler, api.GetTodosConsumer))
	r.Handle("/{id}", httptransport.New(api, api.GetTodoHandler, api.GetTodoConsumer))
	r.Handle("/", httptransport.New(api, api.CreateTodoHandler, api.CreateTodoConsumer))
	r.Handle("/{id}", httptransport.New(api, api.UpdateTodoHandler, api.UpdateTodoConsumer))
	r.Handle("/{id}", httptransport.New(api, api.DeleteTodoHandler, api.DeleteTodoConsumer))

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// swaggerSpec retrieves the swagger spec from a local file
func swaggerSpec() *Spec {
	spec, err := loadSpec("swagger.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	return spec
}
