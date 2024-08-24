package main

import (
	"github.com/LRU_cache/cache"
)

func main() {
	cache := cache.Initialise_Cache(5)
	cache.Get("www.google.com")
	cache.Display()
	cache.Get("www.yahoo.com")
	cache.Display()
	cache.Get("www.gmail.com")
	cache.Display()
	cache.Get("www.chatgpt.com")
	cache.Display()
	cache.Get("www.slackboat.com")
	cache.Display()
	cache.Get("www.yahoo.com")
	cache.Display()
	cache.Get("www.google.com")
	cache.Display()
	cache.Get("www.yahoo.com")
	cache.Display()
	cache.Get("www.youtube.com")
	cache.Display()
	cache.Get("www.google.com")
	cache.Display()

}
