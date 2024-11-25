package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/antonfisher/json-schema-go"
	"github.com/go-json-rest/rest"
	"github.com/xeipuuv/gojsonschema"
)

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	ISBN   string  `json:"isbn"`
	Price  float32 `json:"price"`
	Pages  int     `json:"pages"`
}

func main() {
	// Load the initial schema
	schemaFile, err := ioutil.ReadFile("schema.json")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the schema validator
	loader := gojsonschema.NewBytesLoader(schemaFile)
	schema, err := gojsonschema.NewSchema(loader)
	if err != nil {
		log.Fatal(err)
	}

	// Create a REST API handler
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/books", validateAndUpdateSchema(schema)),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func validateAndUpdateSchema(schema *gojsonschema.Schema) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		defer r.Body.Close()
		var book Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			rest.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validate the data against the schema
		data, err := json.Marshal(book)
		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := schema.Validate(gojsonschema.NewBytesLoader(data))
		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !result.Valid() {
			for _, err := range result.Errors() {
				log.Printf("Validation error: %s\n", err)
			}
			rest.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}
		// Data is valid, update the schema automatically
		err = updateSchema(&book, schema)
		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteJson(&book)
	}
}

func updateSchema(book *Book, schema *gojsonschema.Schema) error {
	// Incremental schema evolution is a complex task.
	// In this simple example, we just remove the old schema file and
	// generate a new one based on the modified `Book` struct.
	// Note that this approach may not be suitable for large and complex schemas.
	if err := os.Remove("schema.json"); err != nil && !os.IsNotExist(err) {
		return err
	}
	schemaFile, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("schema.json", schemaFile, 0644)
}
