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

// Search returns whether a key is in bst
func (t *BinarySearchTree) Search(key int) bool {
	return t.search(t.Root, key)
}

func (t *BinarySearchTree) search(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if root.Data == key {
		return true
	}

	if root.Data > key {
		return t.search(root.Left, key)
	} else {
		return t.search(root.Right, key)
	}
}

// Delete remove a node from and keep the tree struct
func (t *BinarySearchTree) Delete(key int) bool {
	return t.delete(t.Root, key)
}

func (t *BinarySearchTree) delete(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if root.Data == key {
		return t.del(root)
	} else if root.Data > key {
		return t.delete(root.Right, key)
	} else {
		return t.delete(root.Left, key)
	}
}

func (t *BinarySearchTree) del(root *Node) bool {
	if root.Left == nil {
		root = root.Left
	} else if root.Right == nil {
		root = root.Right
	} else {
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
	}
	return true
}

func (t *BinarySearchTree) String() string {
	return t.inOrder(t.Root)
}

func (t *BinarySearchTree) inOrder(root *Node) string {
	s := ""
	if root != nil {
		s += t.inOrder(root.Left)
		s += fmt.Sprintf("%v\t", root.Data)
		s += t.inOrder(root.Right)
	}
	return s
}
