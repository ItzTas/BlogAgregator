package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ItsTas/BlogAgregator/internal/client"
	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DBURL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	cfg := &apiConfig{
		DB:     dbQueries,
		Client: client.NewClient(5*time.Minute, 5*time.Minute),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readiness)
	mux.HandleFunc("GET /v1/err", errorResponse)

	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(cfg.handlerGetUser))

	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.handlerCreateFeed))
	mux.HandleFunc("GET /v1/feeds", cfg.handlerRetrieveFeeds)

	mux.HandleFunc("POST /v1/feed_follows", cfg.middlewareAuth(cfg.handlerCreateFeedFollow))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", cfg.handlerDeleteFeedFollow)
	mux.HandleFunc("GET /v1/feed_follows", cfg.middlewareAuth(cfg.handlerRetrieveFeedFollowFromUser))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Listening on port: %s\n", port)

	log.Fatal(srv.ListenAndServe())
}
