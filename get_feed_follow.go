package main

import (
	"net/http"

	"github.com/ItsTas/BlogAgregator/internal/database"
)

func (cfg *apiConfig) handlerRetrieveFeedFollowFromUser(w http.ResponseWriter, r *http.Request, user database.User) {
	dbFfs, err := cfg.DB.GetFeedFollowsById(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get the feeds follows of the user")
		return
	}

	fls := make([]FeedFollow, len(dbFfs))
	for i, ff := range dbFfs {
		fls[i] = FeedFollow(ff)
	}

	respondWithJSON(w, http.StatusOK, fls)
}
