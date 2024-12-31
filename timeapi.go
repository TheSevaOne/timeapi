package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime  string `json:"current_time,omitempty"`
	ErrorMessage string `json:"err,omitempty"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Query().Get("error") == "true" {
		response := TimeResponse{CurrentTime: "", ErrorMessage: "An error occurred while fetching the time."}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	_time := time.Now().Format(time.RFC822)
	response := TimeResponse{CurrentTime: _time, ErrorMessage: ""}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8081", nil)
}
