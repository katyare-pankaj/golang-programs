package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Property represents a property in the system
type Property struct {
	ID      uint   `gorm:"primary_key"`
	Address string `gorm:"not null"`
	City    string `gorm:"not null"`
	State   string `gorm:"not null"`
	Zip     string `gorm:"not null"`
	Price   uint   `gorm:"not null"`
	Sold    bool   `gorm:"default:false"`
}

// InitDB initializes the database connection
func InitDB() *gorm.DB {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db
}

// CreatePropertyHandler handles the creation of a new property
func CreatePropertyHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation code for creating a new property
}

// GetPropertyHandler handles the retrieval of a property by its ID
func GetPropertyHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation code for retrieving a property by its ID
}

// UpdatePropertyHandler handles the updating of a property by its ID
func UpdatePropertyHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation code for updating a property by its ID
}

// DeletePropertyHandler handles the deletion of a property by its ID
func DeletePropertyHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation code for deleting a property by its ID
}

func main() {
	// Initialize the database
	db := InitDB()
	defer db.Close()

	// Create the table if it doesn't exist
	db.AutoMigrate(&Property{})

	// Create a new router
	router := mux.NewRouter()

	// Register the routes
	router.HandleFunc("/properties", CreatePropertyHandler).Methods("POST")
	router.HandleFunc("/properties/{id}", GetPropertyHandler).Methods("GET")
	router.HandleFunc("/properties/{id}", UpdatePropertyHandler).Methods("PUT")
	router.HandleFunc("/properties/{id}", DeletePropertyHandler).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
