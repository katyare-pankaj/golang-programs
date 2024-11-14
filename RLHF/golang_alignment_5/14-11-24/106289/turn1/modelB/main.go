package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// DummyRiskCalculation simulates a risk calculation function
func DummyRiskCalculation() error {
	time.Sleep(time.Duration(100 * time.Millisecond)) // Simulate some work
	return nil
}

func riskHandler(w http.ResponseWriter, r *http.Request) {
	err := DummyRiskCalculation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Risk calculation completed")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/risk", riskHandler)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
