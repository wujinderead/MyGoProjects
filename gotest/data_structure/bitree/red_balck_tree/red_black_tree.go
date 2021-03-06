package red_balck_tree

import (
	"bytes"
	list2 "container/list"
	"fmt"
)

const (
	Red   = true
	Black = false
)

type RedBlackNode struct {
	left, right, parent *RedBlackNode
	key                 int
	value               interface{}
	color               bool
}

type RedBlackTree struct {
	Root *RedBlackNode
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{nil}
}

func (tree *RedBlackTree) rotateLeft(p *RedBlackNode) {
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

func (tree *RedBlackTree) rotateRight(p *RedBlackNode) {
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

func (node *RedBlackNode) String() string {
	if node.color == Red {
		return fmt.Sprintf("\033[1;31m%d\033[0m", node.key)
	} else {
		return fmt.Sprintf("\033[1;30m%d\033[0m", node.key)
	}
}

func stringifyNode(node *RedBlackNode) string {
	if node == nil {
		return "xx"
	}
	if node.color == Red {
		return fmt.Sprintf("\033[1;31m%02d\033[0m", node.key)
	} else {
		return fmt.Sprintf("\033[1;30m%02d\033[0m", node.key)
	}
}

func (node *RedBlackNode) predecessor() *RedBlackNode {
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

func (node *RedBlackNode) successor() *RedBlackNode {
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

func (tree *RedBlackTree) Set(key int, value interface{}) {
	if tree.Root == nil {
		tree.Root = &RedBlackNode{key: key, value: value, color: Black}
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
				cur.left = &RedBlackNode{key: key, value: value, parent: cur}
				tree.fixAfterInsertion(cur.left)
				return
			}
			cur = cur.left
		}
		if key > cur.key {
			if cur.right == nil {
				cur.right = &RedBlackNode{key: key, value: value, parent: cur}
				tree.fixAfterInsertion(cur.right)
				return
			}
			cur = cur.right
		}
	}
}

func (tree *RedBlackTree) fixAfterInsertion(x *RedBlackNode) {
	x.color = Red // set is any way
	for x != nil && x != tree.Root && colorOf(parentOf(x)) == Red {
		if parentOf(x) == leftOf(parentOf(parentOf(x))) { // p is pp's left son
			sib := rightOf(parentOf(parentOf(x))) // sib is pp's right son
			if colorOf(sib) == Red {
				setColor(sib, Black)
				setColor(parentOf(x), Black)
				setColor(parentOf(parentOf(x)), Red)
				x = parentOf(parentOf(x))
			} else {
				if x == rightOf(parentOf(x)) {
					x = parentOf(x)
					tree.rotateLeft(x) // to this to make x == leftOf(parentOf(x))
				}
				setColor(parentOf(x), Black)
				setColor(parentOf(parentOf(x)), Red)
				tree.rotateRight(parentOf(parentOf(x)))
			}
		} else { // p is pp's right son
			sib := leftOf(parentOf(parentOf(x))) // sib is pp's left son
			if colorOf(sib) == Red {
				setColor(sib, Black)
				setColor(parentOf(x), Black)
				setColor(parentOf(parentOf(x)), Red)
				x = parentOf(parentOf(x))
			} else {
				if x == leftOf(parentOf(x)) {
					x = parentOf(x)
					tree.rotateRight(x)
				}
				setColor(parentOf(x), Black)
				setColor(parentOf(parentOf(x)), Red)
				tree.rotateLeft(parentOf(parentOf(x)))
			}
		}
	}
	setColor(tree.Root, Black)
}

func (tree *RedBlackTree) fixAfterDeletion(x *RedBlackNode) {
	for x != tree.Root && colorOf(x) == Black {
		if x == leftOf(parentOf(x)) {
			sib := rightOf(parentOf(x))
			if colorOf(sib) == Red {
				setColor(sib, Black)
				setColor(parentOf(x), Red)
				tree.rotateLeft(parentOf(x))
				sib = rightOf(parentOf(x))
			}
			if colorOf(leftOf(sib)) == Black && colorOf(rightOf(sib)) == Black {
				setColor(sib, Red)
				x = parentOf(x)
			} else {
				if colorOf(rightOf(sib)) == Black {
					setColor(leftOf(sib), Black)
					setColor(sib, Red)
					tree.rotateRight(sib)
					sib = rightOf(parentOf(x))
				}
				setColor(sib, colorOf(parentOf(x)))
				setColor(parentOf(x), Black)
				setColor(rightOf(sib), Black)
				tree.rotateLeft(parentOf(x))
				x = tree.Root
			}
		} else {
			sib := leftOf(parentOf(x))
			if colorOf(sib) == Red {
				setColor(sib, Black)
				setColor(parentOf(x), Red)
				tree.rotateRight(parentOf(x))
				sib = leftOf(parentOf(x))
			}
			if colorOf(rightOf(sib)) == Black && colorOf(leftOf(sib)) == Black {
				setColor(sib, Red)
				x = parentOf(x)
			} else {
				if colorOf(leftOf(sib)) == Black {
					setColor(rightOf(sib), Black)
					setColor(sib, Red)
					tree.rotateLeft(sib)
					sib = leftOf(parentOf(x))
				}
				setColor(sib, colorOf(parentOf(x)))
				setColor(parentOf(x), Black)
				setColor(leftOf(sib), Black)
				tree.rotateRight(parentOf(x))
				x = tree.Root
			}
		}
	}
	setColor(x, Black)
}

func parentOf(x *RedBlackNode) *RedBlackNode {
	if x != nil {
		return x.parent
	}
	return nil
}

func leftOf(x *RedBlackNode) *RedBlackNode {
	if x != nil {
		return x.left
	}
	return nil
}

func rightOf(x *RedBlackNode) *RedBlackNode {
	if x != nil {
		return x.right
	}
	return nil
}

func colorOf(x *RedBlackNode) bool {
	if x != nil {
		return x.color
	}
	return Black // color for nil is black
}

func setColor(x *RedBlackNode, color bool) {
	if x != nil {
		x.color = color
	}
}

func (tree *RedBlackTree) Get(key int) interface{} {
	node := tree.getNode(key)
	if node != nil {
		return node.value
	}
	return nil
}

func (tree *RedBlackTree) getNode(key int) *RedBlackNode {
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

func (tree *RedBlackTree) Remove(key int) interface{} {
	node := tree.getNode(key)
	if node == nil {
		return nil
	}
	v := node.value
	if node.left != nil && node.right != nil { // both sons are non-nil
		successor := node.successor() // successor must exist
		node.key = successor.key
		node.value = successor.value // copy successor's kv to node
		node = successor             // and remove successor
	}
	var son *RedBlackNode = nil
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
			node.parent.right = son // use son to replace node
		}
		node.left, node.right, node.parent = nil, nil, nil
		if node.color == Black { // only fix deletion when it's black
			tree.fixAfterDeletion(son)
		}
	} else if node.parent == nil { // deleting root, and root has no son
		tree.Root = nil
	} else { // both son is nil
		if node.color == Black { // only fix deletion when it's black
			tree.fixAfterDeletion(node)
		}
		if node.parent != nil { // unlink node after fix
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

func (tree *RedBlackTree) String() string {
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
		node, _ := value.(*RedBlackNode)
		if node != nil {
			l.PushBack(node.left)
			l.PushBack(node.right)
			buf.WriteString(node.String() + ", ")
			allNil = false
		} else {
			l.PushBack(nil)
			l.PushBack(nil)
			buf.WriteString("x, ")
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

func (tree *RedBlackTree) Print() {
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
		node, _ := value.(*RedBlackNode)
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

func (tree *RedBlackTree) Traverse(eachNode func(node *RedBlackNode)) {
	l := list2.New()
	if tree.Root != nil {
		l.PushBack(tree.Root)
	}
	for l.Len() > 0 {
		value := l.Remove(l.Front())
		node, _ := value.(*RedBlackNode)
		eachNode(node)
		if node.left != nil {
			l.PushBack(node.left)
		}
		if node.right != nil {
			l.PushBack(node.right)
		}
	}
}

type iter struct {
	cur *RedBlackNode
}

func (tree *RedBlackTree) Iter() *iter {
	first := tree.Root
	if first == nil {
		return &iter{nil}
	}
	for first.left != nil {
		first = first.left
	}
	return &iter{first}
}

func (tree *RedBlackTree) ReverseIter() *iter {
	last := tree.Root
	if last == nil {
		return &iter{nil}
	}
	for last.left != nil {
		last = last.right
	}
	return &iter{last}
}

func (it *iter) HasPrev() bool {
	return it.cur != nil
}

func (it *iter) Prev() *RedBlackNode {
	cur := it.cur
	it.cur = cur.predecessor()
	return cur
}

func (it *iter) HasNext() bool {
	return it.cur != nil
}

func (it *iter) Next() *RedBlackNode {
	cur := it.cur
	it.cur = cur.successor()
	return cur
}
