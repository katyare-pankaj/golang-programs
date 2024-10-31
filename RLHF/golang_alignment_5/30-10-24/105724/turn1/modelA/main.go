// main.go
package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn1/modelA/handler"
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn1/modelA/logging"
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn1/modelA/processor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	logger := logging.NewLogger()
	dataProcessor := processor.NewDataProcessor()

	router := mux.NewRouter()
	router.HandleFunc("/process", handler.NewProcessHandler(logger, dataProcessor)).Methods("POST")

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
