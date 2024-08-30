package lru_cache

import "fmt"

type cache struct {
	hashmap    map[string]*node
	queue      *queue
	max_length int32
}

func Initialise_Cache(max_length int32) *cache {
	if max_length < 2 {
		panic("The Cache Size Can Not Less Than 2..........")
	}
	return &cache{
		hashmap:    make(map[string]*node),
		queue:      newQueue(),
		max_length: max_length,
	}
}

func (c *cache) Get(key string) string {
	node, ok := c.hashmap[key]
	if ok {
		c.queue.addExistingToBegining(node)
	} else {

		if c.queue.length == c.max_length {
			rmv_node := c.queue.tail
			c.queue.removeNode(rmv_node)
			delete(c.hashmap, rmv_node.val)
		}

		node = c.queue.addNewValueToBegining(key)
		c.hashmap[key] = node
	}
	return node.val
}

func (c *cache) Display() {
	node := c.queue.head
	var i int32
	for i = 0; i < c.queue.length; i++ {
		fmt.Printf("{%s}-->", node.val)
		node = node.right
	}
	fmt.Println()
}
