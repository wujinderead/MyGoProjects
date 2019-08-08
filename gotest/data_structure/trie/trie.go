package trie

type Trie struct {
	Root *TrieNode
}

const dictSize = 26

type TrieNode struct {
	children [dictSize]*TrieNode
	field    interface{}
}

func NewTrie() Trie {
	root := new(TrieNode)
	return Trie{root}
}

func (trie Trie) Put(key string, value interface{}) {
	if value == nil {
		panic("cannot put nil value")
	}
	checkKey(key)
	cur := trie.Root
	for _, v := range key {
		if cur.children[v-'a'] == nil {
			cur.children[v-'a'] = new(TrieNode)
		}
		cur = cur.children[v-'a']
	}
	cur.field = value
}

func (trie Trie) Contain(key string) bool {
	_, v := trie.Get(key)
	return v != nil
}

func (trie Trie) Get(key string) (*TrieNode, interface{}) {
	checkKey(key)
	cur := trie.Root
	for _, v := range key {
		if cur.children[v-'a'] == nil {
			return nil, nil
		}
		cur = cur.children[v-'a']
	}
	return cur, cur.field
}

func checkKey(key string) {
	for _, v := range key {
		if v > 'z' || v < 'a' {
			panic("only lowercase alphabet supported")
		}
	}
}
