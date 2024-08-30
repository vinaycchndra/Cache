package lru_cache

type node struct {
	val   string
	left  *node
	right *node
}

func newNode() *node {
	node := node{}
	return &node
}

func (n *node) setValue(str string) {
	n.val = str
}
