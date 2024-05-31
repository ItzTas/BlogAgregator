package main

import (
	"errors"
	"net/http"

	"github.com/ItsTas/BlogAgregator/internal/auth"
	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	})
}

func (cfg *apiConfig) getUserFromRequest(r *http.Request) (database.User, error) {
	apiKey, err := auth.ParseAPIKey(r.Header)
	if err != nil {
		return database.User{}, errors.New("invalid apiKey")
	}
	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil || user.ID == uuid.Nil {
		return database.User{}, errors.New("could not find the user")
	}
	return user, nil
}
