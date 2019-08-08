package avl_tree

import (
	list2 "container/list"
	"fmt"
)

type AvlNode struct {
	left, right *AvlNode
	key         int
	value       interface{}
	height      int
}

type AvlTree struct {
	Root  *AvlNode
	stack *stack
}

func NewAvlTree() *AvlTree {
	return &AvlTree{Root: nil, stack: newStack()}
}

func (tree *AvlTree) rotateLeft(p *AvlNode) *AvlNode {
	if p == nil || p.right == nil {
		return nil
	}
	r := p.right
	p.right = r.left
	r.left = p
	p.height = max(heightOf(p.left), heightOf(p.right)) + 1
	r.height = max(heightOf(r.left), heightOf(r.right)) + 1
	if p == tree.Root {
		tree.Root = r
	}
	return r
}

func (tree *AvlTree) rotateRight(p *AvlNode) *AvlNode {
	if p == nil || p.left == nil {
		return nil
	}
	l := p.left
	p.left = l.right
	l.right = p
	p.height = max(heightOf(p.left), heightOf(p.right)) + 1
	l.height = max(heightOf(l.left), heightOf(l.right)) + 1
	if p == tree.Root {
		tree.Root = l
	}
	return l
}

func (node *AvlNode) String() string {
	//return fmt.Sprintf("[%d,%d]", node.key, node.height)
	return fmt.Sprintf("[%2d]", node.key)
}

func heightOf(node *AvlNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (tree *AvlTree) Set(key int, value interface{}) {
	if tree.Root == nil {
		tree.Root = &AvlNode{key: key, value: value, height: 1}
		return
	}
	cur := tree.Root
	tree.stack.reinit()
	for {
		if cur.key == key {
			cur.value = value
			return // no new node added, return
		}
		if key < cur.key {
			if cur.left == nil {
				cur.left = &AvlNode{key: key, value: value, height: 1}
				cur.height = 2
				tree.fixAfterInsert(tree.stack, cur, cur.left)
				return
			}
			tree.stack.push(cur)
			cur = cur.left
		}
		if key > cur.key {
			if cur.right == nil {
				cur.right = &AvlNode{key: key, value: value, height: 1}
				cur.height = 2
				tree.fixAfterInsert(tree.stack, cur, cur.right)
				return
			}
			tree.stack.push(cur)
			cur = cur.right
		}
	}
}

func (tree *AvlTree) fixAfterInsert(stack *stack, p, q *AvlNode) {
	var r, newr *AvlNode
	for stack.len() > 0 {
		r = stack.pop()
		diff := heightOf(r.left) - heightOf(r.right)
		if diff >= -1 && diff <= 1 { // r balanced
			r.height = max(heightOf(r.left), heightOf(r.right)) + 1 // update height
			q = p
			p = r
			continue
		}
		// unbalanced, rotate
		if p == r.left {
			if q == p.right { // left right
				r.left = tree.rotateLeft(p)
			}
			// left left
			newr = tree.rotateRight(r) // store new r
		} else {
			if q == p.left { // right left
				r.right = tree.rotateRight(p)
			}
			// right right
			newr = tree.rotateLeft(r) // store new r
		}
		if stack.len() > 0 { // peek r's parent to set pointer
			rp := stack.peek()
			if rp.left == r {
				rp.left = newr
			} else {
				rp.right = newr
			}
		}
		break // only fix for once
	}
}

func stringifyNode(node *AvlNode) string {
	if node == nil {
		return "xx"
	}
	return fmt.Sprintf("%2d", node.key)
}

func (tree *AvlTree) Remove(key int) interface{} {
	cur := tree.Root
	tree.stack.reinit()
	for cur != nil && cur.key != key {
		tree.stack.push(cur)
		if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if cur == nil { // if value not found
		return nil
	}
	// found node, store its value
	value := cur.value
	// if both son non-nil, delete successor
	if cur.left != nil && cur.right != nil {
		tree.stack.push(cur)
		p := cur.right
		for p.left != nil {
			tree.stack.push(p)
			p = p.left
		}
		cur.key = p.key
		cur.value = p.value
		cur = p // set cur to successor
	}
	var p *AvlNode = nil
	if tree.stack.len() > 0 {
		p = tree.stack.peek()
	}
	var son *AvlNode = nil
	if cur.left != nil {
		son = cur.left
	} else {
		son = cur.right
	}
	if son != nil { // one son non-nil
		if p == nil { // deleting root, and root has a son
			tree.Root = son
		} else if p.left == cur {
			p.left = son
		} else {
			p.right = son
		}
		cur.left, cur.right = nil, nil
		// fix
		tree.fixAfterDelete(tree.stack)
	} else if p == nil { // deleting root, root has no son
		tree.Root = nil
	} else { // deleting a leaf
		if p.left == cur {
			p.left = nil
		} else {
			p.right = nil
		}
		// fix
		tree.fixAfterDelete(tree.stack)
	}
	return value
}

// stack stores the deleted node's parents.
// delete a node can immediately make its parent unbalance,
// so process the parents one by one.
func (tree *AvlTree) fixAfterDelete(stack *stack) {
	for stack.len() > 0 {
		cur := stack.pop()
		oldHeight := cur.height
		newcur := cur
		leftHeight := heightOf(cur.left)
		rightHeight := heightOf(cur.right)
		// fix unbalance
		if leftHeight-rightHeight > 1 {
			l := cur.left
			if heightOf(l.right) > heightOf(l.left) { // left right
				cur.left = tree.rotateLeft(l)
			} // left left
			newcur = tree.rotateRight(cur)
		} else if rightHeight-leftHeight > 1 {
			r := cur.right
			if heightOf(r.left) > heightOf(r.right) { // right left
				cur.right = tree.rotateRight(r)
			} // right right
			newcur = tree.rotateLeft(cur)
		}

		// if rebalance preformed, newcur's height is already ok
		newcur.height = max(heightOf(newcur.left), heightOf(newcur.right)) + 1
		if stack.len() > 0 { // peek cur's parent to set pointer
			parent := stack.peek()
			if parent.left == cur {
				parent.left = newcur
			} else {
				parent.right = newcur
			}
		}
		// no matter if we have fixed unbalance ot not,
		// if the current height unchanged, no need for further fix
		if newcur.height == oldHeight {
			break
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

func (tree *AvlTree) Print() {
	if tree.Root == nil {
		fmt.Println("[nil]")
		return
	}
	space := "  "
	l := list2.New()
	l.PushBack(tree.Root)
	lenTier := 1
	curTier := 0
	allNil := true
	tiers := make([][]string, 0)
	tier := make([]string, lenTier)
	for {
		value := l.Remove(l.Front())
		node, _ := value.(*AvlNode)
		tier[curTier] = stringifyNode(node)
		curTier++
		if node != nil {
			l.PushBack(node.left)
			l.PushBack(node.right)
			allNil = false
		} else {
			l.PushBack(nil)
			l.PushBack(nil)
		}
		if curTier == lenTier {
			if allNil {
				break
			}
			curTier = 0
			lenTier = 2 * lenTier
			tiers = append(tiers, tier)
			tier = make([]string, lenTier)
			allNil = true
		}
	}
	lenTier = len(tiers)
	for i, t := range tiers {
		init := (1 << uint(lenTier-i-1)) - 1
		inter := (1 << uint(lenTier-i)) - 1
		for i := 0; i < init; i++ {
			fmt.Print(space)
		}
		for index, ele := range t {
			fmt.Print(ele)
			if index < len(t)-1 {
				for i := 0; i < inter; i++ {
					fmt.Print(space)
				}
			}
		}
		fmt.Println()
	}
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type stack struct {
	top   int
	slice []*AvlNode
}

func newStack() *stack {
	return &stack{-1, make([]*AvlNode, 64)}
}

func (s *stack) push(node *AvlNode) {
	s.top++
	s.slice[s.top] = node
}

func (s *stack) pop() *AvlNode {
	s.top--
	return s.slice[s.top+1]
}

func (s *stack) peek() *AvlNode {
	return s.slice[s.top]
}

func (s *stack) len() int {
	return s.top + 1
}

func (s *stack) reinit() {
	s.top = -1
}
