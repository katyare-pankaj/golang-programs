package main

// Import required packages
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a Database struct
type Database struct {
	DB *sql.DB
}

// Define a Global variable to store the database connections
var DBs map[string]*Database

// Initialize the databases
func initDBs() {
	DBs = make(map[string]*Database)
	// Initialize connections to multiple databases (example: DBs["db1"], DBs["db2"], etc.)
	// Open the database connections and store them in the DBs map
}

// Handle the transaction logic
func handleTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dbName := vars["dbname"] // Get the database name from the request URL

	// Check if the requested database exists
	db, ok := DBs[dbName]
	if !ok {
		http.Error(w, "Database not found", http.StatusNotFound)
		return
	}

	// Start a new transaction
	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback() // Rollback the transaction if there's an error

	// Perform the database operations within the transaction
	result, err := tx.Exec("INSERT INTO table1 (value) VALUES (?)", "some value")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the transaction was successful
	rowAffected, err := result.RowsAffected()
	if err != nil || rowAffected == 0 {
		http.Error(w, "Transaction failed", http.StatusInternalServerError)
		return
	}

	// Commit the transaction if all operations are successful
	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	fmt.Fprintf(w, "Transaction successful!")
}

func main() {
	// Initialize the databases
	initDBs()

	// Create a new router
	r := mux.NewRouter()

	// Define the route for handling transactions
	r.HandleFunc("/transaction/{dbname}", handleTransaction).Methods("POST")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", r))
}
