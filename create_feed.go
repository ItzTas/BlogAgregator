package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	var lastFetchedAt *time.Time
	if feed.LastFetchedAt.Valid {
		lastFetchedAt = &feed.LastFetchedAt.Time
	} else {
		lastFetchedAt = nil
	}

	newFeed := Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAt: lastFetchedAt,
	}
	return newFeed
}

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramethers struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	type returnVals struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
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

	newFeed, err := cfg.DB.CreateFeed(r.Context(), feedParams)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Feed already exists")
		return
	}

	ff, err := cfg.newfeedFollow(newFeed.ID, user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, returnVals{
		Feed:       databaseFeedToFeed(newFeed),
		FeedFollow: ff,
	})
}
