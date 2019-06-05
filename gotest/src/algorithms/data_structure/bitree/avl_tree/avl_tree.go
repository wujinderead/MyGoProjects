package avl_tree

import (
	"bytes"
	list2 "container/list"
	"fmt"
)

type AvlNode struct {
	left, right, parent *AvlNode
	key                 int
	value               interface{}
}

type AvlTree struct {
	Root *AvlNode
}

func NewAvlTree() *AvlTree {
	return &AvlTree{}
}

func (tree *AvlTree) rotateLeft(p *AvlNode) {
	if p == nil || p.right == nil {
		return
	}
	r := p.right
	p.right = r.left
	if r.left != nil {
		r.left.parent = p
	}
	r.parent = p.parent
	if p.parent == nil {
		tree.Root = r
	} else if p.parent.right == p {
		p.parent.right = r
	} else {
		p.parent.left = r
	}
	r.left = p
	p.parent = r
}

func (tree *AvlTree) rotateRight(p *AvlNode) {
	if p == nil || p.left == nil {
		return
	}
	l := p.left
	p.left = l.right
	if l.right != nil {
		l.right.parent = p
	}
	l.parent = p.parent
	if p.parent == nil {
		tree.Root = l
	} else if p.parent.right == p {
		p.parent.right = l
	} else {
		p.parent.left = l
	}
	l.right = p
	p.parent = l
}

func (node *AvlNode) String() string {
	return fmt.Sprintf("[key=%d,value=%v]", node.key, node.value)
}

func (node *AvlNode) predecessor() *AvlNode {
	if node.left != nil {
		node := node.left
		for node.right != nil {
			node = node.right
		}
		return node
	}
	p := node.parent
	for p != nil {
		if node == p.right {
			return p
		}
		node = p
		p = p.parent
	}
	return nil
}

func (node *AvlNode) successor() *AvlNode {
	if node.right != nil {
		node := node.right
		for node.left != nil {
			node = node.left
		}
		return node
	}
	p := node.parent
	for p != nil {
		if node == p.left {
			return p
		}
		node = p
		p = p.parent
	}
	return nil
}

func (tree *AvlTree) Set(key int, value interface{}) {
	if tree.Root == nil {
		tree.Root = &AvlNode{key: key, value: value}
		return
	}
	cur := tree.Root
	for {
		if cur.key == key {
			cur.value = value
			return
		}
		if key < cur.key {
			if cur.left == nil {
				cur.left = &AvlNode{key: key, value: value, parent: cur}
				return
			}
			cur = cur.left
		}
		if key > cur.key {
			if cur.right == nil {
				cur.right = &AvlNode{key: key, value: value, parent: cur}
				return
			}
			cur = cur.right
		}
	}
}

func (tree *AvlTree) Get(key int) interface{} {
	node := tree.getNode(key)
	if node != nil {
		return node.value
	}
	return nil
}

func (tree *AvlTree) getNode(key int) *AvlNode {
	if tree.Root == nil {
		return nil
	}
	cur := tree.Root
	for cur != nil {
		if cur.key == key {
			return cur
		}
		if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}

func (tree *AvlTree) Remove(key int) interface{} {
	node := tree.getNode(key)
	if node == nil {
		return nil
	}
	v := node.value
	if node.left != nil && node.right != nil { // both sons are non-nil
		successor := node.successor()
		node.key = successor.key
		node.value = successor.value
		node = successor
	}
	var son *AvlNode = nil
	if node.left != nil {
		son = node.left
	} else {
		son = node.right
	}
	if son != nil { // only one son is non-nil
		son.parent = node.parent
		if node.parent == nil {
			tree.Root = son
		} else if node == node.parent.left {
			node.parent.left = son
		} else {
			node.parent.right = son
		}
		node.left, node.right, node.parent = nil, nil, nil
	} else if node.parent == nil { // deleting root, and root has no son
		tree.Root = nil
	} else { // both son is nil
		if node.parent != nil {
			if node == node.parent.left {
				node.parent.left = nil
			} else {
				node.parent.right = nil
			}
			node.parent = nil
		}
	}
	return v
}

func (tree *AvlTree) String() string {
	if tree.Root == nil {
		return "[nil]"
	}
	l := list2.New()
	l.PushBack(tree.Root)
	lenTier := 1
	curTier := 0
	allNil := true
	buf := new(bytes.Buffer)
	curSize := 0
	for {
		value := l.Remove(l.Front())
		curTier++
		node, _ := value.(*AvlNode)
		if node != nil {
			l.PushBack(node.left)
			l.PushBack(node.right)
			buf.WriteString(node.String() + ", ")
			allNil = false
		} else {
			l.PushBack(nil)
			l.PushBack(nil)
			buf.WriteString("[nil], ")
		}
		if curTier == lenTier {
			if allNil {
				break
			}
			buf.WriteString("\n")
			curTier = 0
			lenTier = 2 * lenTier
			allNil = true
			curSize = buf.Len()
		}
	}
	buf.Truncate(curSize - 1)
	return buf.String()
}

func (tree *AvlTree) Traverse(eachNode func(node *AvlNode)) {
	l := list2.New()
	if tree.Root != nil {
		l.PushBack(tree.Root)
	}
	for l.Len() > 0 {
		value := l.Remove(l.Front())
		node, _ := value.(*AvlNode)
		eachNode(node)
		if node.left != nil {
			l.PushBack(node.left)
		}
		if node.right != nil {
			l.PushBack(node.right)
		}
	}
}
