package lfu_cache

import (
	"errors"
)

// type Item any

// func (i *Item) lessThan(j *Item) bool {

// }

type Item interface {
	lessThan(*Item) bool
}

type minHeap struct {
	max_heap_size int32
	current_size  int32
	heap          []*Item
}

func newMinHeap(max_heap_size int32) *minHeap {
	if max_heap_size == int32(0) {
		max_heap_size = int32(1)
	}
	heap := make([]*Item, 0, max_heap_size)
	return &minHeap{max_heap_size: max_heap_size, current_size: 0, heap: heap}
}

// Heapifying the slice
func (h *minHeap) heapify(index int) {
	i := int32(index)

	for i < h.Len() {
		object := h.heap[i]
		new_index := i

		// Comparing the object at the ith index with the left child.
		left := i*int32(2) + int32(1)
		if left < h.Len() {

			left_obj := *(h.heap)[left]
			if left_obj.lessThan(object) {
				new_index = left
				object = &left_obj
			}
		}

		// Comparing the object at the ith index or left child with the right child.
		right := left + int32(1)
		if right < h.Len() {
			right_obj := *(h.heap)[right]
			if right_obj.lessThan(object) {
				new_index = right
			}
		}

		// Swaping the objects if the heap is still not in order.
		if i != new_index {
			h.swap(i, new_index)
			i = new_index
		} else {
			// Breaking from the loop since the heap is in order.
			break
		}

	}
}

func (h *minHeap) heapPop() *Item {
	if h.Len() == 0 {
		return nil
	}
	// Swapping the last and first element since first element needs to be popped.
	index := h.Len() - int32(1)

	item := h.heap[0]

	h.heap[0] = h.heap[index]

	// Deleting Last Element Since it is now at the beginning.
	h.heap[index] = nil
	h.current_size--

	// Heapifying the slice since it is not in minHeap order
	h.heapify(0)
	return item
}

func (h *minHeap) heapPush(item *Item) error {
	if h.Len() == h.max_heap_size {
		return errors.New("Heap Size Over Flow Can Not Add New Item")
	}

	// Assigning the new item at the last index
	index := h.Len()
	h.heap[index] = item
	h.current_size++
	current_item := *item

	for index > 0 {
		parent_index := (index - int32(1)) / int32(2)
		parent := h.heap[parent_index]

		if current_item.lessThan(parent) {
			h.swap(index, parent_index)
			index = parent_index
		} else {
			break
		}
	}
	return nil
}

func (h *minHeap) swap(i, j int32) {
	temp := h.heap[i]
	h.heap[i] = h.heap[j]
	h.heap[j] = temp
}

func (h *minHeap) Len() int32 {
	return h.current_size
}
