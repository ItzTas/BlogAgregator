package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type paramethers struct {
		Name string `name:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramethers{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode params")
	}
}
