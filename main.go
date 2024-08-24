package main

import (
	"github.com/LRU_cache/cache"
)

func main() {
	cache := cache.Initialise_Cache(10)
	cache.Get("www.google.com")
	cache.Display()
	cache.Get("www.yahoo.com")
	cache.Display()
	cache.Get("www.google.com")
	cache.Display()
	cache.Get("www.yahoo.com")
	cache.Display()
	cache.Get("www.google.com")
	cache.Display()
	cache.Get("www.yahoo.com")
	cache.Display()
	cache.Get("www.google.com")
	cache.Display()
	cache.Get("www.yahoo.com")
	cache.Display()
}
