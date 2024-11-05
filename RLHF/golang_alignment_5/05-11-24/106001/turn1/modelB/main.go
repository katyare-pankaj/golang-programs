package main

import (
	"go-programs/RLHF/golang_alignment_5/05-11-24/106001/turn1/modelB/api"
	"log"
	"net/http"
)

func main() {
	http.Handle("/api/weather", api.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
