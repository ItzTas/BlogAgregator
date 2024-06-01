package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request) {
	fdID := r.PathValue("feedFollowID")
	if fdID == "" {
		respondWithJSON(w, http.StatusBadRequest, "Did not specify and id")
		return
	}

	uuidfdID, err := uuid.Parse(fdID)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Could not parse the id")
		return
	}

	_, err = cfg.DB.DeleteFeedFollow(r.Context(), uuidfdID)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Could not delete the feed follow")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
