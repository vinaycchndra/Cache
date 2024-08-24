-Least Recently Used Cache in Golang. A least recently used cache stores the order of items in a most recently used order or the items which are not accessed very recently will be pushed to the last in the queue.

- We will use a doubly linked list data structure to mantain the order of the cached values.
                                                               
- We will use hash map to access the cached value in O(1) time.

- We will use another object(struct) to access the head and tail of the doubly linked list and matains the maximum value of the value of the cache value.


Process: 
-We will initialize the Cache having the Queue and hashMap. The Que will be initialized with the head and tail with the nil values. These only two nodes also initialize our doubly linked 	list.

- When we fetch a value from the cache the value becomes the recently use it should we removed and should be added to the begining of the linked list and if we exceeded the length of the cache the node value at the last should be removed since it is least fetched.

- If we are accessing the values which are not precent in the doubly linked list will not alter the structure of the cache.

- If we are adding anything to the cache that value should also be added at the begining of the queue.
