/*
Package trie implement a trie datasruct:
	In computer science, a **trie**, also called digital tree and sometimes
	radix tree or prefix tree (as they can be searched by prefixes),
	is a kind of search tree—an ordered tree data structure that is
	used to store a dynamic set or associative array where the keys
	are usually strings. Unlike a binary search tree, no node in the
	tree stores the key associated with that node; instead, its
	position in the tree defines the key with which it is associated.
	All the descendants of a node have a common prefix of the string
	associated with that node, and the root is associated with the
	empty string. Values are not necessarily associated with every
	node. Rather, values tend only to be associated with leaves,
	and with some inner nodes that correspond to keys of interest.
	For the space-optimized presentation of prefix tree, see compact
	prefix tree.
WikiPage:
	* https://en.wikipedia.org/wiki/Trie
*/
package trie

// ALPHAS means alphas nums
const ALPHAS = 26

type node struct {
	// isKey use to mark whether this node is the end of a word.
	// isKey bool
	// count use to do word-frequency count, it is a substitute of isKey.
	count    int
	children [ALPHAS]*node
}

// Trie means aa datastructure of trie
type Trie struct {
	root *node
}

// New is Trie`s constructor
func New() *Trie {
	return &Trie{
		root: &node{},
	}
}

// Insert inserts a word to the trie
func (t *Trie) Insert(w string) {
	cur := t.root
	for _, c := range []byte(w) {
		i := c - 'a'
		if cur.children[i] == nil {
			cur.children[i] = &node{}
		}
		cur = cur.children[i]
	}
	cur.count++
}

// Search a word from the trie
func (t *Trie) Search(w string) int {
	cur := t.root
	for _, c := range []byte(w) {
		i := c - 'a'
		if cur.children[i] == nil {
			return 0
		}
		cur = cur.children[i]
	}
	return cur.count
}
