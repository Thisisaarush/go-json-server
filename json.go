package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func resondWithError(w http.ResponseWriter, r *http.Request, code int, message string) {
	if code > 499 {
		log.Println("Responding with error: ", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, r, code, errResponse{Error: message})
}