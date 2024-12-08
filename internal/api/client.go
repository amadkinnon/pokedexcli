package api

import (
	"net/http"
	"pokedexcli/internal/cache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      cache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
