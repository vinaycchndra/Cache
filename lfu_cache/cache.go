package lfu_cache

import "fmt"

type cache struct {
	hashmap         map[string]*Item
	min_heap        *minHeap
	curr_time_stamp int32
	max_length      int32
}

func Initialise_Cache(max_length int32) *cache {
	return &cache{
		hashmap:         make(map[string]*Item),
		min_heap:        newMinHeap(max_length),
		curr_time_stamp: int32(-1),
		max_length:      max_length,
	}
}

func (c *cache) Get(key string) any {
	item, ok := c.hashmap[key]
	if ok {
		c.curr_time_stamp++
		value := item.getValue()
		item.increaseCount()
		item.setTime(c.curr_time_stamp)
		c.min_heap.heapify(item.getIndex())
		return value
	}
	return nil
}

func (c *cache) Set(key string, value any) {
	item, ok := c.hashmap[key]
	if ok {
		c.curr_time_stamp++
		item.setValue(value)
		item.increaseCount()
		item.setTime(c.curr_time_stamp)
		c.min_heap.heapify(item.getIndex())
	} else {
		c.curr_time_stamp++
		new_item := newItem(key, value, int32(1), c.curr_time_stamp)
		err := c.min_heap.heapPush(new_item)
		if err != nil {
			// Getting error means the cache is full now so we need to remove the evicted item,
			// manually ourself from the hashmap and the heap.
			removed_item := c.min_heap.heapPop()
			delete(c.hashmap, removed_item.getKey())
			c.min_heap.heapPush(new_item)
		}
		c.hashmap[new_item.getKey()] = new_item
	}
}

func (c *cache) PrintArr() {
	for _, value := range c.min_heap.heap {
		if value != nil {
			fmt.Print(*value, " ")
		}
	}
	fmt.Println()
}
