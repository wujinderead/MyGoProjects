package trie

import (
	"container/list"
	"fmt"
	"strings"
	"testing"
)

func TestLenString(t *testing.T) {
	s1 := "abc"
	fmt.Println(len(s1)) // 3
	s2 := "千里驰援有天霞"
	fmt.Println(len(s2)) // 21, 21 bytes for utf-8
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i]) // 3 bytes
	}
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i]) // 21 bytes
	}
	for _, v := range s1 {
		fmt.Println(v) // 3 runes
	}
	for _, v := range s2 {
		fmt.Println(v) // 7 runes
	}
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Put("aaa", 1)
	trie.Put("abc", 2)
	trie.Put("aab", 3)
	trie.Put("b", 4)
	trie.Put("", 5)
	fmt.Println(trie.Get("aaa"))
	fmt.Println(trie.Get("aa"))
	fmt.Println(trie.Get("abc"))
	fmt.Println(trie.Get("abb"))
	fmt.Println(trie.Get("b"))
	fmt.Println(trie.Get(""))
}

var test = "there are several ways to represent tries corresponding to different tradeoffs between memory use and speed of the operations the basic form is that of a linked set of nodes where each node contains an array of child pointers one for each symbol in the alphabet so for the english alphabet one would store  child pointers and for the alphabet of bytes  pointers this is simple but wasteful in terms of memory using the alphabet of bytes size  and fourbyte pointers each node requires a kilobyte of storage and when there is little overlap in the strings prefixes the number of required nodes is roughly the combined length of the stored strings put another way the nodes near the bottom of the tree tend to have few children and there are many of them so the structure wastes space storing null pointersthe storage problem can be alleviated by an implementation technique called alphabet reduction whereby the original strings are reinterpreted as longer strings over a smaller alphabet eg a string of n bytes can alternatively be regarded as a string of n fourbit units and stored in a trie with sixteen pointers per node lookups need to visit twice as many nodes in the worst case but the storage requirements go down by a factor of eightan alternative implementation represents a node as a triple symbol child next and links the children of a node together as a singly linked list child points to the nodes first child next to the parent nodes next child the set of children can also be represented as a binary search tree one instance of this idea is the ternary search tree developed by bentley and sedgewick another alternative in order to avoid the use of an array of  pointers ascii as suggested before is to store the alphabet array as a bitmap of  bits representing the ascii alphabet reducing dramatically the size of the nodes bitwise tries this section does not cite any sources please help improve this section by adding citations to reliable sources unsourced material may be challenged and removed february  learn how and when to remove this template message bitwise tries are much the same as a normal character based trie except that individual bits are used to traverse what effectively becomes a form of binary tree generally implementations use a special cpu instruction to very quickly find the first set bit in a fixed length key eg gccs builtinclz intrinsic this value is then used to index a  or entry table which points to the first item in the bitwise trie with that number of leading zero bits the search then proceeds by testing each subsequent bit in the key and choosing child or child appropriately until the item is found although this process might sound slow it is very cache local and highly parallelizable due to the lack of register dependencies and therefore in fact has excellent performance on modern outoforder execution cpus a red black tree for example performs much better on paper but is highly cache unfriendly and causes multiple pipeline and tlb stalls on modern cpus which makes that algorithm bound by memory latency rather than cpu speed in comparison a bitwise trie rarely accesses memory and when it does it does so only to read thus avoiding smp cache coherency overhead hence it is increasingly becoming the algorithm of choice for code that performs many rapid insertions and deletions such as memory allocators eg recent versions of the famous doug leas allocator dlmalloc and its descendents compressing tries compressing the trie and merging the common branches can sometimes yield large performance gains this works best under the following conditionsthe trie is mostly static key insertions to or deletions from a prefilled trie are disabled citation needed only lookups are needed the trie nodes are not keyed by node specific data or the nodes data are common the total set of stored keys is very sparse within their representation space citation neededfor example it may be used to represent sparse bitsets ie subsets of a much larger fixed enumerable set in such a case the trie is keyed by the bit element position within the full set the key is created from the string of bits needed to encode the integral position of each element such tries have a very degenerate form with many missing branches after detecting the repetition of common patterns or filling the unused gaps the unique leaf nodes bit strings can be stored and compressed easily reducing the overall size of the triesuch compression is also used in the implementation of the various fast lookup tables for retrieving unicode character properties these could include casemapping tables eg for the greek letter pi from  to  or lookup tables normalizing the combination of base and combining characters like the aumlaut in german  or the dalet patah dagesh ole in biblical hebrew  for such applications the representation is similar to transforming a very large unidimensional sparse table eg unicode code points into a multi dimensional matrix of their combinations and then using the coordinates in the hypermatrix as the string key of an uncompressed trie to represent the resulting character the compression will then consist of detecting and merging the common columns within the hypermatrix to compress the last dimension in the key for example to avoid storing the full multibyte unicode code point of each element forming a matrix column the groupings of similar code points can be exploited each dimension of the hypermatrix stores the start position of the next dimension so that only the offset typically a single byte need be stored the resulting vector is itself compressible when it is also sparse so each dimension associated to a layer level in the trie can be compressed separately some implementations do support such data compression within dynamic sparse tries and allow insertions and deletions in compressed tries however this usually has a significant cost when compressed segments need to be split or merged some tradeoff has to be made between data compression and update speed a typical strategy is to limit the range of global lookups for comparing the common branches in the sparse trie citation needed the result of such compression may look similar to trying to transform the trie into a directed acyclic graph dag because the reverse transform from a dag to a trie is obvious and always possible however the shape of the dag is determined by the form of the key chosen to index the nodes in turn constraining the compression possible another compression strategy is to unravel the data structure into a single byte array this approach eliminates the need for node pointers substantially reducing the memory requirements this in turn permits memory mapping and the use of virtual memory to efficiently load the data from diskone more approach is to pack the trie liang describes a spaceefficient implementation of a sparse packed trie applied to automatic hyphenation in which the descendants of each node may be interleaved in memory external memory tries several trie variants are suitable for maintaining sets of strings in external memory including suffix trees a combination of trie and btree called the btrie has also been suggested for this task compared to suffix trees they are limited in the supported operations but also more compact while performing update operations faster"

func TestTrieMap(t *testing.T) {
	strs := strings.Split(test, " ")
	maper := make(map[string]int)
	for _, str := range strs {
		if n, ok := maper[str]; ok {
			maper[str] = n + 1
		} else {
			maper[str] = 1
		}
	}
	trie := NewTrie()
	for _, str := range strs {
		face, value := trie.Get(str)
		if face != nil {
			if value != nil {
				face.field = value.(int) + 1
			} else {
				face.field = 1
			}
		} else {
			trie.Put(str, 1)
		}
	}
	for k, v := range maper {
		_, tv := trie.Get(k)
		if tv == nil || tv.(int) != v {
			fmt.Println("===", k)
			t.Fatal("get map wrong")
		}
		//fmt.Println(k, "=", v, tv.(int))
	}

	// breadth first search to display trie
	type node struct {
		tnode *TrieNode
		tier  int
		char  string
	}
	ll := list.New()
	al, _ := trie.Get("al")
	ll.PushBack(&node{al, 0, "al"})
	curTier := 0
	for ll.Len() > 0 {
		front := ll.Front()
		ll.Remove(front)
		cur := front.Value.(*node)
		if cur.tier > curTier {
			curTier = cur.tier
			fmt.Println()
		}
		if cur.tnode.field != nil {
			fmt.Print(cur.char, " ", cur.tnode.field, ", ")
		} else {
			fmt.Print(cur.char, ", ")
		}
		for i, child := range cur.tnode.children {
			if child != nil {
				ll.PushBack(&node{child, curTier + 1, cur.char + string(rune(i+'a'))})
			}
		}
	}
	fmt.Println()
}
