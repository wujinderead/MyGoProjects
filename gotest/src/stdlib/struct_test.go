package stdlib

import (
	"fmt"
	"testing"
)

type node struct {
	value int
}

type tree struct {
	root *node
}

type rbnode struct {
	node
	color bool
}

type rbtree struct {
	root *rbnode
}

func (node *node) str() {
	fmt.Println(node.value)
}

func (tree *tree) str() {
	tree.root.str()
}

func TestStructInherit(tt *testing.T) {
	n := &node{10}
	t := &tree{n}
	n.str()
	t.str()
	rbn := &rbnode{node{10}, true}
	//rbt := &rbtree{rbn}
	rbn.str()
}
