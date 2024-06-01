package main

import (
	"net/http"

	"github.com/ItsTas/BlogAgregator/internal/database"
)

func (cfg *apiConfig) handlerRetrieveFeedFollowFromUser(w http.ResponseWriter, r *http.Request, user database.User) {
	dbFds, err := cfg.DB.GetFeedFollowsById(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get the feeds follows of the user")
	}

	fds := make([]FeedFollow, len(dbFds))
	for i, fd := range dbFds {
		fds[i] = FeedFollow(fd)
	}

	respondWithJSON(w, http.StatusOK, fds)
}
