package pokeapi

import (
	"net/http"
	"internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client {
		httpClient: http.Client {
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) StopCacheReap() {
	c.cache.Done <- true
}