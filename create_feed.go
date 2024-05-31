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
	type returnVals struct {
		Feed       Feed
		FeedFollow FeedFollow
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
		respondWithError(w, http.StatusBadRequest, "Feed already exists")
		return
	}

	fd, err := cfg.newfeedFollow(f.ID, user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, returnVals{
		Feed:       Feed(feedParams),
		FeedFollow: fd,
	})
}
