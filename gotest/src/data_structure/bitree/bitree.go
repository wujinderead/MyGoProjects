package bitree

import (
	"bytes"
	list2 "container/list"
	"fmt"
)

type BiTreeNode struct {
	left, right, parent *BiTreeNode
	key                 int
	value               interface{}
}

type BiTree struct {
	Root *BiTreeNode
}

func NewBiTree() *BiTree {
	return &BiTree{}
}

func (tree *BiTree) rotateLeft(p *BiTreeNode) {
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

func (tree *BiTree) rotateRight(p *BiTreeNode) {
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

func (node *BiTreeNode) String() string {
	return fmt.Sprintf("[key=%d,value=%v]", node.key, node.value)
}

func (node *BiTreeNode) predecessor() *BiTreeNode {
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

func (node *BiTreeNode) successor() *BiTreeNode {
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

func (tree *BiTree) Set(key int, value interface{}) {
	if tree.Root == nil {
		tree.Root = &BiTreeNode{key: key, value: value}
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
				cur.left = &BiTreeNode{key: key, value: value, parent: cur}
				return
			}
			cur = cur.left
		}
		if key > cur.key {
			if cur.right == nil {
				cur.right = &BiTreeNode{key: key, value: value, parent: cur}
				return
			}
			cur = cur.right
		}
	}
}

func (tree *BiTree) Get(key int) interface{} {
	node := tree.getNode(key)
	if node != nil {
		return node.value
	}
	return nil
}

func (tree *BiTree) getNode(key int) *BiTreeNode {
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

func (tree *BiTree) Remove(key int) interface{} {
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
	var son *BiTreeNode = nil
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

func (tree *BiTree) String() string {
	if tree.Root == nil {
		return "[nil]"
	}
	l := list2.New()
	l.PushBack(tree.Root)
	len_tier := 1
	cur_tier := 0
	all_nil := true
	buf := new(bytes.Buffer)
	cur_size := 0
	for {
		value := l.Remove(l.Front())
		cur_tier++
		node, _ := value.(*BiTreeNode)
		if node != nil {
			l.PushBack(node.left)
			l.PushBack(node.right)
			buf.WriteString(node.String() + ", ")
			all_nil = false
		} else {
			l.PushBack(nil)
			l.PushBack(nil)
			buf.WriteString("[nil], ")
		}
		if cur_tier == len_tier {
			if all_nil {
				break
			}
			buf.WriteString("\n")
			cur_tier = 0
			len_tier = 2 * len_tier
			all_nil = true
			cur_size = buf.Len()
		}
	}
	buf.Truncate(cur_size - 1)
	return buf.String()
}

func (tree *BiTree) Traverse(eachNode func(node *BiTreeNode)) {
	l := list2.New()
	if tree.Root != nil {
		l.PushBack(tree.Root)
	}
	for l.Len() > 0 {
		value := l.Remove(l.Front())
		node, _ := value.(*BiTreeNode)
		eachNode(node)
		if node.left != nil {
			l.PushBack(node.left)
		}
		if node.right != nil {
			l.PushBack(node.right)
		}
	}
}
