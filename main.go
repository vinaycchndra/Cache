package main

import (
	"github.com/LRU_cache/cache"
)

func main() {
	node := cache.NewNode("Hello We are ready to Go ! with a brand new start value")
	node.Display()
}
