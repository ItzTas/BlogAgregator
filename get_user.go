package main

import (
	"errors"
	"net/http"

	"github.com/ItsTas/BlogAgregator/internal/auth"
	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
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

func databaseUserToUser(user database.User) User {
	return User(user)
}
