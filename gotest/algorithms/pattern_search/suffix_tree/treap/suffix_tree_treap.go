package treap

import (
	"container/list"
	"math/rand"
	"time"
)

// for large character set for example utf8, it has much more characters than ascii.
// so it's not practical to use array 'children [alphabet_size]*SuffixTreeNode'. instead,
// a variety of data structure can be used, which have different time and space costs.
//                        |  Lookup     |  Insert     |  Travel
// -----------------------+-------------+-------------+------------
//   sibling list         |  O(σ)       |  Θ(1)       |  Θ(1)
//   hashmap              |  Θ(1)       |  Θ(1)       |  O(σ)
//   balanced search tree |  O(logσ)    |  O(logσ)    |  O(1)
//   sorted array         |  O(logσ)    |  O(σ)       |  O(1)
//   hashmap+sibling list |  O(1)       |  O(1)       |  O(1)
//
// base on: https://en.wikipedia.org/wiki/Suffix_tree

var (
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type SuffixTreeNode struct {
	children    *treapRoot
	suffixLink  *SuffixTreeNode
	start       int
	end         *int // use pointer to make all ends expand with O(1) time
	suffixIndex int
}

type SuffixTree struct {
	Root  *SuffixTreeNode
	Text  string
	Runes []rune
}

func NewSuffixTreeTreap(text string) *SuffixTree {
	root := new(SuffixTreeNode)
	activeNode := root
	activeEdgeIndex := -1
	activeLength := 0
	end := new(int)
	remainCount := 0
	runes := []rune(text)
	for i := 0; i <= len(runes); i++ { // i, current index
		// get current character
		var curRune rune
		if i == len(runes) {
			curRune = 0
		} else {
			curRune = runes[i]
		}

		*end = i                                  // increment all end
		remainCount++                             // increment remain suffix count and loop
		var preInternalNode *SuffixTreeNode = nil // for suffix link

		for remainCount > 0 {
			// walk down to the active point
			// mistake 4
			for activeLength > 0 {
				// if active length>0, there must be an edge
				curEdge := activeNode.getChild(runes[activeEdgeIndex])
				curEdgeLen := *curEdge.end - curEdge.start + 1
				// if active len == edge len (can only be equal)
				// walk down to reset active node
				if activeLength >= curEdgeLen {
					activeLength -= curEdgeLen
					activeEdgeIndex += curEdgeLen
					activeNode = curEdge
				} else {
					break
				}
			}

			// active length=0, check if need to add a new leaf
			// active length>0, check if need to split an edge and make an new internal node
			if activeLength == 0 {
				// if active length==0, set active edge to current character
				activeEdgeIndex = i

				// check if active edge going out of active node
				if activeNode.getChild(curRune) == nil {
					// active edge not present, create a new edge
					newNode := new(SuffixTreeNode)
					newNode.start = i
					newNode.end = end
					activeNode.setChild(curRune, newNode)
					// create new node (add new suffix), need to decrease remainCount
					remainCount--

					// if previous internal node not null, link suffix link
					if preInternalNode != nil {
						preInternalNode.suffixLink = activeNode
						// set preInternalNode to null to prevent unexpected link
						// mistake 1
						preInternalNode = nil
					}
				} else {
					// active edge present, suffix won't be added explicitly in current phase
					// increment active length and exit current phase
					activeLength = 1
					// mistake 5
					if preInternalNode != nil {
						preInternalNode.suffixLink = activeNode
					}
					break
				}
			} else {
				// if active length>0, check whether current text character present after active point
				curEdge := activeNode.getChild(runes[activeEdgeIndex])
				activePointIndex := curEdge.start + activeLength - 1
				if runes[activePointIndex+1] == curRune {
					// current text character present after active point,
					// suffix won't be added explicitly in current phase
					// increment active length and exit current phase
					activeLength++
					break
				} else {
					// mistake 2
					// split current edge, the trick here is:
					// use current edge to contain remain characters,
					// create a new node as current node's father and active node's child
					newNode := new(SuffixTreeNode)
					newNode.start = curEdge.start
					newNode.end = new(int)
					*newNode.end = activePointIndex
					curEdge.start = activePointIndex + 1 // modify start and end
					// newNode become activeNode's child
					activeNode.setChild(runes[activeEdgeIndex], newNode)
					// curEdge (the edge to split) become newNode's child
					newNode.setChild(runes[activePointIndex+1], curEdge)

					// create new leaf for current text character
					newLeaf := new(SuffixTreeNode)
					newLeaf.start = i
					newLeaf.end = end                  // leaf end equals to global end
					newNode.setChild(curRune, newLeaf) // leaf added to new node

					// if previous internal node not null, link suffix link
					if preInternalNode != nil {
						preInternalNode.suffixLink = newNode
					}
					preInternalNode = newNode

					// add new suffix, need to decrease remainCount
					remainCount--
				}
			}

			// find next active point
			if activeNode == root && activeLength > 0 {
				// if active node is root, next active node is still root
				// just change active edge and decrement active length
				// if activeLength=0, it must be the last added suffix,
				// i.e., remainCount=0, thus no need to find next
				activeEdgeIndex = i - remainCount + 1
				activeLength--
			} else if activeNode != root {
				// mistake 3
				// if active node is internal node, use suffix link
				// no need to decrement active length when use suffix link
				activeNode = activeNode.suffixLink
			}
		}
	}

	dfsToSetSuffixIndex(root)
	return &SuffixTree{Root: root, Text: text, Runes: runes}
}

func dfsToSetSuffixIndex(root *SuffixTreeNode) {
	curLen := 0
	stack := list.New()
	it := root.children.Iter()
	for it.hasNext() {
		stack.PushBack(it.next().value)
	}
	visited := make(map[*SuffixTreeNode]struct{})
	for stack.Len() > 0 {
		cur := stack.Back().Value.(*SuffixTreeNode)
		if _, ok := visited[cur]; !ok { // not visited, peek and add children
			visited[cur] = struct{}{}
			curLen += *cur.end - cur.start + 1
			if cur.children != nil { // internal node
				it = cur.children.Iter()
				for it.hasNext() {
					stack.PushBack(it.next().value)
				}
			} else { // leaf
				cur.suffixIndex = *cur.end - curLen + 1
			}
		} else { // visited, pop
			stack.Remove(stack.Back())
			curLen -= *cur.end - cur.start + 1
		}
	}
}

func (node *SuffixTreeNode) getChild(r rune) *SuffixTreeNode {
	if node.children == nil {
		return nil
	}
	t := node.children.get(r)
	if t != nil {
		return t.value
	}
	return nil
}

func (node *SuffixTreeNode) setChild(r rune, v *SuffixTreeNode) {
	if node.children == nil {
		node.children = &treapRoot{}
	}
	node.children.set(r, v)
}

type treapRoot struct {
	root *treap
}

type treap struct {
	left, right, parent *treap
	ticket              int32
	key                 rune
	value               *SuffixTreeNode
}

func (tr *treapRoot) leftRotate(t *treap) {
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
		tr.root = r
	} else if t.parent.left == t {
		t.parent.left = r
	} else {
		t.parent.right = r
	}
	r.parent = t.parent
	t.parent = r
}

func (tr *treapRoot) rightRotate(t *treap) {
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
		tr.root = l
	} else if t.parent.left == t {
		t.parent.left = l
	} else {
		t.parent.right = l
	}
	l.parent = t.parent
	t.parent = l
}

func (tr *treapRoot) get(key rune) *treap {
	t := tr.root
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

func (tr *treapRoot) set(key rune, v *SuffixTreeNode) {
	if tr.root == nil {
		tr.root = &treap{nil, nil, nil, rander.Int31(), key, v}
		return
	}
	t := tr.root
	for {
		if key == t.key {
			t.value = v
			return
		} else if key < t.key {
			if t.left == nil {
				t.left = &treap{nil, nil, t, rander.Int31(), key, v}
				tr.fixAfterInsertion(t.left)
				return
			}
			t = t.left
		} else {
			if t.right == nil {
				t.right = &treap{nil, nil, t, rander.Int31(), key, v}
				tr.fixAfterInsertion(t.right)
				return
			}
			t = t.right
		}
	}
}

func (tr *treapRoot) fixAfterInsertion(t *treap) {
	// maintain a big heap
	// check if there is a violation of big heap until t==root or no violation
	for t != tr.root && t.parent.ticket < t.ticket {
		if t.parent.left == t {
			tr.rightRotate(t.parent)
		} else {
			tr.leftRotate(t.parent)
		}
	}
}

func (t *treap) successor() *treap {
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

func (t *treap) predecessor() *treap {
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
	cur *treap
}

func (tr *treapRoot) Iter() *iter {
	first := tr.root
	if first == nil {
		return &iter{first}
	}
	for first.left != nil {
		first = first.left
	}
	return &iter{first}
}

func (tr *treapRoot) ReverseIter() *iter {
	last := tr.root
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

func (it *iter) prev() *treap {
	cur := it.cur
	it.cur = cur.predecessor()
	return cur
}

func (it *iter) hasNext() bool {
	return it.cur != nil
}

func (it *iter) next() *treap {
	cur := it.cur
	it.cur = cur.successor()
	return cur
}
