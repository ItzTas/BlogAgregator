package main

import (
	"net/http"

	"github.com/ItsTas/BlogAgregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := cfg.getUserFromRequest(r)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		handler(w, r, user)
	}
}
