// main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Initialize database
	db, err := gorm.Open("sqlite3", "social.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Migrate database tables
	db.AutoMigrate(&User{}, &Post{})

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/posts", postsHandler).Methods("GET")
	r.HandleFunc("/posts/create", createPostHandler).Methods("POST")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// homeHandler handles the home page request
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Social Media Application!")
}

// postsHandler handles the posts page request
func postsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Posts page!")
}

// createPostHandler handles the create post request
func createPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create post page!")
}

// User represents a user in the social media application
type User struct {
	gorm.Model
	Name string
}

// Post represents a post in the social media application
type Post struct {
	gorm.Model
	Content string
	UserID  uint
}
