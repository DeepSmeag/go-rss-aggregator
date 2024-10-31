package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		// internal errors
		log.Println("Server 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"` // this instructs the json.Marshal function to put this field under the "error" key, so error: <Error>

	}
	respondWithJSON(w, code, errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshal JSON response: %v", payload)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
