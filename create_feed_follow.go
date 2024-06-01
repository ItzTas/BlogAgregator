package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramethers struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramethers{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode paramethers")
		return
	}

	fd, err := cfg.newfeedFollow(params.FeedID, user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, fd)
}

func (cfg *apiConfig) newfeedFollow(feedID uuid.UUID, user database.User) (FeedFollow, error) {
	uid := uuid.New()
	fdParams := database.CreateFeedFollowParams{
		ID:        uid,
		FeedID:    feedID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	fd, err := cfg.DB.CreateFeedFollow(context.TODO(), fdParams)
	if err != nil {
		return FeedFollow{}, errors.New("feed follow already exists")
	}
	return FeedFollow(fd), nil
}
