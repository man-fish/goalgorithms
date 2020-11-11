/*
Package binarysearchtree implements a binary search tree:
	In computer science, **binary search trees** (BST), sometimes called
	ordered or sorted binary trees, are a particular type of container:
	data structures that store "items" (such as numbers, names etc.)
	in memory. They allow fast lookup, addition and removal of
	items, and can be used to implement either dynamic sets of
	items, or lookup tables that allow finding an item by its v
	(e.g., finding the phone number of a person by name).

	Binary search trees keep their vs in sorted order, so that lookup
	and other operations can use the principle of binary search:
	when looking for a v in a tree (or a place to insert a new v),
	they traverse the tree from root to leaf, making comparisons to
	vs stored in the nodes of the tree and deciding, on the basis
	of the comparison, to continue searching in the left or right
	subtrees. On average, this means that each comparison allows
	the operations to skip about half of the tree, so that each
	lookup, insertion or deletion takes time proportional to the
	logarithm of the number of items stored in the tree. This is
	much better than the linear time required to find items by v
	in an (unsorted) array, but slower than the corresponding
	operations on hash tables.
BinarySearchTree on Wiki:
	* https://en.wikipedia.org/wiki/Binary_search_tree
*/
package binarysearchtree

import "fmt"

// Node is element of bst
type Node struct {
	Data        int
	Left, Right *Node
}

// BinarySearchTree represent a binary search tree
type BinarySearchTree struct {
	Root *Node
}

// New is a constructor
func New(data []int) *BinarySearchTree {
	if len(data) < 1 {
		return nil
	}
	bst := new(BinarySearchTree)
	bst.Root = &Node{Data: data[0]}
	for i := 1; i < len(data); i++ {
		bst.insert(bst.Root, data[i])
	}
	return bst
}

// Insert an element to bst
func (t *BinarySearchTree) Insert(i int) {
	if t.Root == nil {
		t.Root = &Node{Data: i}
		return
	}
	t.insert(t.Root, i)
}

func (t *BinarySearchTree) insert(node *Node, i int) {
	if i > node.Data {
		if node.Right == nil {
			node.Right = &Node{Data: i}
		} else {
			t.insert(node.Right, i)
		}
	} else {
		if node.Left == nil {
			node.Left = &Node{Data: i}
		} else {
			t.insert(node.Left, i)
		}
	}
}

// Search returns whether a v is in bst
func (t *BinarySearchTree) Search(v int) bool {
	return t.search(t.Root, v)
}

func (t *BinarySearchTree) search(root *Node, v int) bool {
	if root == nil {
		return false
	}

	if root.Data == v {
		return true
	}

	if root.Data > v {
		return t.search(root.Left, v)
	} else {
		return t.search(root.Right, v)
	}
}

// Delete remove a node from and keep the tree struct
func (t *BinarySearchTree) Delete(v int) {
	t.Root = t.delete(t.Root, v)
}

func (t *BinarySearchTree) delete(root *Node, v int) *Node {
	if root == nil {
		return nil
	}

	if root.Data == v {
		if root.Left != nil && root.Right != nil {
			cur := root
			next := root.Left

			for next.Right != nil {
				cur = next
				next = next.Right
			}

			root.Data = next.Data

			if cur == root {
				cur.Left = next.Left
			} else {
				cur.Right = next.Left
			}
		} else {
			if root.Left != nil {
				root = root.Left
			} else {
				root = root.Right
			}
		}
		return root
	} else if root.Data > v {
		root.Left = t.delete(root.Left, v)
	} else {
		root.Right = t.delete(root.Right, v)
	}
	return root
}

func (t *BinarySearchTree) String() string {
	return t.inOrder(t.Root)
}

func (t *BinarySearchTree) inOrder(root *Node) string {
	s := ""
	if root != nil {
		s += t.inOrder(root.Left)
		s += fmt.Sprintf("%v,", root.Data)
		s += t.inOrder(root.Right)
	}
	return s
}
