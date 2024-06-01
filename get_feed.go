package main

import "net/http"

func (cfg *apiConfig) handlerRetrieveFeeds(w http.ResponseWriter, r *http.Request) {
	dbfeeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get the feeds")
		return
	}

	feeds := make([]Feed, len(dbfeeds))
	for i, feed := range dbfeeds {
		feeds[i] = databaseFeedToFeed(feed)
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
