package cache

type queue struct {
	head   *node
	tail   *node
	length int32
}

func newQueue() *queue {
	queue := queue{head: newNode(), tail: newNode(), length: 0}
	return &queue
}

func (q *queue) removeNode(node *node) {
	if node == q.tail {
		node.left.right = nil

	} else if node != q.head {
		node.left.right = node.right
		node.right.left = node.left
	}
}

func (q *queue) addExistingToBegining(node *node) {
	q.removeNode(node)
	q.head = node
}

func (q *queue) addNewValueToBegining(key string) *node {
	node := newNode()
	node.setValue(key)
	node.right = q.head
	q.head.left = node
	q.head = node
	q.length++
	return node
}
