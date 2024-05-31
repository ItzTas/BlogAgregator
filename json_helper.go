package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		fmt.Printf("Responding with 5XX code: %v, msg: %v", code, msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}
