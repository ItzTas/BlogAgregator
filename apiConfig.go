package main

import (
	"github.com/ItsTas/BlogAgregator/internal/client"
	"github.com/ItsTas/BlogAgregator/internal/database"
)

type apiConfig struct {
	DB     *database.Queries
	Client client.Client
}
