package treap

import (
	"math/rand"
	"time"
)

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

type Treap struct {
	Root *TreapNode
}

func NewTreap() *Treap {
	return &Treap{nil}
}

type TreapNode struct {
	left, right, parent *TreapNode
	ticket              int64
	key                 int
	value               interface{}
}

func (tr *Treap) leftRotate(t *TreapNode) {
	if t == nil || t.right == nil {
		return
	}
	r := t.right
	rl := r.left
	r.left = t
	t.right = rl
	if rl != nil {
		rl.parent = t
	}
	if t.parent == nil {
		tr.Root = r
	} else if t.parent.left == t {
		t.parent.left = r
	} else {
		t.parent.right = r
	}
	r.parent = t.parent
	t.parent = r
}

func (tr *Treap) rightRotate(t *TreapNode) {
	if t == nil || t.left == nil {
		return
	}
	l := t.left
	lr := l.right
	l.right = t
	t.left = lr
	if lr != nil {
		lr.parent = t
	}
	if t.parent == nil {
		tr.Root = l
	} else if t.parent.left == t {
		t.parent.left = l
	} else {
		t.parent.right = l
	}
	l.parent = t.parent
	t.parent = l
}

func (tr *Treap) Get(key int) *TreapNode {
	t := tr.Root
	for t != nil {
		if key == t.key {
			return t
		} else if key < t.key {
			t = t.left
		} else {
			t = t.right
		}
	}
	return nil
}

func (tr *Treap) Set(key int, v interface{}) {
	if tr.Root == nil {
		tr.Root = &TreapNode{nil, nil, nil, rander.Int63(), key, v}
		return
	}
	t := tr.Root
	for {
		if key == t.key {
			t.value = v
			return
		} else if key < t.key {
			if t.left == nil {
				t.left = &TreapNode{nil, nil, t, rander.Int63(), key, v}
				tr.fixAfterInsertion(t.left)
				return
			}
			t = t.left
		} else {
			if t.right == nil {
				t.right = &TreapNode{nil, nil, t, rander.Int63(), key, v}
				tr.fixAfterInsertion(t.right)
				return
			}
			t = t.right
		}
	}
}

func (tr *Treap) Remove(key int) interface{} {
	node := tr.Get(key)
	if node == nil {
		return nil
	}
	v := node.value
	// both null, check ticket of both child.
	// if left child's ticket is larger, rotate right on node
	// if right child's ticket is larger, rotate left on node
	for node.left != nil && node.right != nil {
		if node.left.ticket > node.right.ticket {
			tr.rightRotate(node)
		} else {
			tr.leftRotate(node)
		}
	}
	// now node has only one son, or nil son
	var son *TreapNode = nil
	if node.left != nil {
		son = node.left
	} else {
		son = node.right
	}
	if son != nil { // has one son
		if node.parent == nil { // delete root, son becomes root
			tr.Root = son
		} else if node.parent.left == node {
			node.parent.left = son
		} else {
			node.parent.right = son
		}
		node.left, node.right, node.parent = nil, nil, nil
	} else if node.parent == nil { // we are deleting root and root has no child, the tree is cleared
		tr.Root = nil
	} else { // both son nil, leaf
		if node.parent.left == node { // unlink parent
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		node.parent = nil
	}
	return v
}

func (tr *Treap) fixAfterInsertion(t *TreapNode) {
	// maintain a big heap
	// check if there is a violation of big heap until t==root or no violation
	for t != tr.Root && t.parent.ticket < t.ticket {
		if t.parent.left == t {
			tr.rightRotate(t.parent)
		} else {
			tr.leftRotate(t.parent)
		}
	}
}

func (t *TreapNode) successor() *TreapNode {
	if t.right != nil {
		t = t.right
		for t.left != nil {
			t = t.left
		}
		return t
	}
	p := t.parent
	for p != nil {
		if p.left == t {
			return p
		}
		t = p
		p = p.parent
	}
	return nil
}

func (t *TreapNode) predecessor() *TreapNode {
	if t.left != nil {
		t = t.left
		for t.right != nil {
			t = t.right
		}
		return t
	}
	p := t.parent
	for p != nil {
		if p.right == t {
			return p
		}
		t = p
		p = p.parent
	}
	return nil
}

type iter struct {
	cur *TreapNode
}

func (tr *Treap) Iter() *iter {
	first := tr.Root
	if first == nil {
		return &iter{first}
	}
	for first.left != nil {
		first = first.left
	}
	return &iter{first}
}

func (tr *Treap) ReverseIter() *iter {
	last := tr.Root
	if last == nil {
		return &iter{last}
	}
	for last.right != nil {
		last = last.right
	}
	return &iter{last}
}

func (it *iter) hasPrev() bool {
	return it.cur != nil
}

func (it *iter) prev() *TreapNode {
	cur := it.cur
	it.cur = cur.predecessor()
	return cur
}

func (it *iter) hasNext() bool {
	return it.cur != nil
}

func (it *iter) next() *TreapNode {
	cur := it.cur
	it.cur = cur.successor()
	return cur
}
