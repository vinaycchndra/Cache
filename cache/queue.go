package cache

type queue struct {
	head   *node
	tail   *node
	length int32
}

func newQueue() *queue {
	head_init := newNode()
	tail_init := newNode()
	head_init.right = tail_init
	tail_init.left = head_init
	queue := queue{head: head_init, tail: tail_init, length: 2}
	return &queue
}

func (q *queue) removeNode(node *node) {

	if node == q.tail {
		node.left.right = nil
		q.length--

	} else if node != q.head {
		node.left.right = node.right
		node.right.left = node.left
		q.length--
	}
}

func (q *queue) addExistingToBegining(node *node) {
	q.removeNode(node)
	q.length++
	q.head.left = node
	node.right = q.head
	node.left = nil
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
