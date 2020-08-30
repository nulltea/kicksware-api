package rest

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
)

func newCacheClient() *cache.Client {
	memcached, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.MFU),
		memory.AdapterWithCapacity(100000),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(memcached),
		cache.ClientWithTTL(10 * time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cacheClient
}

func (h *handler) CacheController(next http.Handler) http.Handler {
	return h.cache.Middleware(next)
}