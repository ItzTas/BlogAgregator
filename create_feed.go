package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramethers struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramethers{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode paramethers")
		return
	}

	uid := uuid.New()
	feedParams := database.CreateFeedParams{
		ID:        uid,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	}

	f, err := cfg.DB.CreateFeed(r.Context(), feedParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create feed")
		return
	}
	respondWithJSON(w, http.StatusCreated, Feed{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Url:       f.Url,
		UserID:    f.UserID,
	})
}
