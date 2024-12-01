package pokeapi

import (
	"net/http"
	"time"

	"github.com/wkeebs/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	Cache      pokecache.Cache
	HttpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}
