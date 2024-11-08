package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Property struct
type Property struct {
	ID      uint   `gorm:"primary_key"`
	Address string `gorm:"not null"`
	City    string `gorm:"not null"`
	// Other fields omitted for brevity
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/properties", CreatePropertyHandler).Methods("POST")
	router.HandleFunc("/properties", GetAllPropertiesHandler).Methods("GET")
	router.HandleFunc("/properties/{id}", GetPropertyHandler).Methods("GET")
	router.HandleFunc("/properties/{id}", UpdatePropertyHandler).Methods("PUT")
	router.HandleFunc("/properties/{id}", DeletePropertyHandler).Methods("DELETE")
	router.HandleFunc("/properties/search", SearchPropertiesHandler).Methods("GET")

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Additional handlers for flexible operations

func SearchPropertiesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	city := query.Get("city")
	minPrice := query.Get("min_price")
	maxPrice := query.Get("max_price")

	var properties []Property
	db := getDB()
	defer db.Close()

	query := db.Model(&Property{})

	if city != "" {
		query = query.Where("city = ?", city)
	}

	if minPrice != "" {
		minPriceInt, err := strconv.Atoi(minPrice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		query = query.Where("price >= ?", minPriceInt)
	}

	if maxPrice != "" {
		maxPriceInt, err := strconv.Atoi(maxPrice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		query = query.Where("price <= ?", maxPriceInt)
	}

	err := query.Find(&properties).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the filtered properties
}
