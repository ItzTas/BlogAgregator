package main

import (
	"net/http"
	"strconv"

	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/thoas/go-funk"
)

const defaultLimit = 10

func (cfg *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limitstr := r.URL.Query().Get("limitstr")
	if limitstr == "" {
		limitstr = strconv.Itoa(defaultLimit)
	}

	limit, err := strconv.Atoi(limitstr)
	if err != nil {
		limit = defaultLimit
	}
	getParams := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	dbposts, err := cfg.DB.GetPostsByUser(r.Context(), getParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get the user posts")
		return
	}

	posts := funk.Map(dbposts, func(dbpost database.Post) Post {
		return databasePostToPost(dbpost)
	}).([]Post)

	respondWithJSON(w, http.StatusOK, posts)

}
