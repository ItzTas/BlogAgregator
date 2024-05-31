package main

import "net/http"

func (cfg *apiConfig) handlerRetrieveFeeds(w http.ResponseWriter, r *http.Request) {
	dbfeeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get the feeds")
		return
	}
	feeds := []Feed{}
	for _, feed := range dbfeeds {
		feeds = append(feeds, Feed(feed))
	}
	respondWithJSON(w, http.StatusOK, feeds)
}
