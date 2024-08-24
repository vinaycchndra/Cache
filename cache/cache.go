package cache

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type queue struct {
	head   *Node
	tail   *Node
	length int32
}

type Cache struct {
	hashmap map[string]*Node
	queue   *queue
}

func NewCache() *Cache {
	return &Cache{
		hashmap: make(map[string]*Node),
		queue:   NewQueue(),
	}
}

func NewQueue() *queue {
	queue := queue{head: NewNode(), tail: NewNode(), length: 0}
	return &queue
}

func NewNode() *Node {
	node := Node{
		Left:  nil,
		Right: nil,
	}
	return &node
}
