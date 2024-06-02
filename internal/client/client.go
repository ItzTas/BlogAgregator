package client

import (
	"net/http"
	"time"

	"github.com/ItsTas/BlogAgregator/internal/cache"
)

type Client struct {
	client http.Client
	cache  cache.Cache
}

func NewClient(expiration, timeout time.Duration) Client {
	c := Client{
		client: http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(expiration),
	}
	return c
}
