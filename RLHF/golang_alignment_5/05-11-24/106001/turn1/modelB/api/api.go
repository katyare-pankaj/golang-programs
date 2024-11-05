package api

import (
	"encoding/json"
	"go-programs/RLHF/golang_alignment_5/05-11-24/106001/turn1/modelB/data"
	"net/http"
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		weather := data.GetWeather()
		jsonResponse, err := json.Marshal(weather)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})
}
