package main

import (
	"net/http"

	"github.com/ItsTas/BlogAgregator/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.ParseAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid apiKey")
		return
	}
	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil || user.ID == uuid.Nil {
		respondWithError(w, http.StatusNotFound, "Could not find the user")
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	})
}
