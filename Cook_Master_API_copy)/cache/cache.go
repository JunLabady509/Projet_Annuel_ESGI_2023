package cache

import (
	"net/http"
	"strings"
	"sync"
)

type response struct {
	header http.Header
	code   int
	body   []byte
}

type memCache struct {
	lock sync.RWMutex // protects cache when reading and writing concurrently to it
	data map[string]response
}

var (
	cache = memCache{data: make(map[string]response)}
)

func set(resource string, response *response) {
	cache.lock.Lock() // protect cache when writing to it
	defer cache.lock.Unlock()
	if response == nil {
		delete(cache.data, resource)
	} else {
		cache.data[resource] = *response
	}
}

func get(resource string) *response {
	cache.lock.RLock() // protect cache when reading from it
	defer cache.lock.RUnlock()
	if response, ok := cache.data[resource]; ok {
		return &response
	}
	return nil
}

func copyHeader(src, dst http.Header) {
	for key, list := range src {
		for _, value := range list {
			dst.Add(key, value)
		}
	}
}

// MakeResource returns a string representation of the request URL
func MakeResource(r *http.Request) string {
	if r == nil {
		return ""
	}
	return strings.TrimSuffix(r.URL.RequestURI(), "/") // remove trailing slash
}

// Clean removes all entries from the cache
func Clean() {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	cache.data = map[string]response{}
}

// Drop removes a specific entry from the cache
func Drop(resource string) {
	set(resource, nil)
}

// This function returns true if a cached response xas successfully served
func Serve(w http.ResponseWriter, r *http.Request) bool {
	if w == nil || r == nil {
		return false
	}
	if r.Header.Get("Cache-Control") == "no-cache" {
		return false
	}
	response := get(MakeResource(r))
	if response == nil {
		return false
	}
	copyHeader(response.header, w.Header())
	w.WriteHeader(response.code)
	if r.Method != http.MethodHead {
		w.Write(response.body)
	}
	return true
}
