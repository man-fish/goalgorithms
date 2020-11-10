package avltree

// Node represent a avl tree node
type Node struct {
	Value  int
	height int
	Left   *Node
	Right  *Node
}

// GetBF returns |n.left.height - n.right.height|
func (n *Node) GetBF() int {
	lh, rh := 0, 0
	if n.Left != nil {
		lh = getHeight(n.Left)
	}
	if n.Right != nil {
		lh = getHeight(n.Right)
	}
	bf := lh - rh
	if bf < 0 {
		bf *= -1
	}
	return bf
}

// getHeight returns a node`s height
func getHeight(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

// LRotate left rotate the tree root to balance the tree on RR condition
func LRotate(n *Node) *Node {
	r := n.Right
	n.Right = r.Left
	r.Left = n
	n.height = max(getHeight(n.Left), getHeight(n.Right)) + 1
	r.height = max(getHeight(r.Left), getHeight(r.Right)) + 1
	return r
}

// RRotate right rotate the tree root to balance the tree on LL condition
func RRotate(n *Node) *Node {
	l := n.Left
	n.Left = l.Right
	l.Right = n
	n.height = max(getHeight(n.Left), getHeight(n.Right)) + 1
	l.height = max(getHeight(l.Left), getHeight(l.Right)) + 1
	return l
}

// RLRotate right rotate the right son of root firstly,
// and then left rotate the root to balance the tree on RL condition
func RLRotate(n *Node) *Node {
	n.Right = RRotate(n.Right)
	return LRotate(n)
}

// LRRotate left rotate the right son of root firstly,
// and then right rotate the root to balance the tree on LR condition
func LRRotate(n *Node) *Node {
	n.Left = LRotate(n.Left)
	return RRotate(n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// AVLTree is a balance search tree
type AVLTree struct {
	root *Node
}

// New is a constructor of AVLTree
func New() *AVLTree {
	return new(AVLTree)
}

// Insert insert a node to tree
func (t *AVLTree) Insert(v int) {
	t.root = insert(t.root, v)
}

func insert(root *Node, v int) *Node {
	if root == nil {
		root = new(Node)
		root.Value = v
	} else if v < root.Value {
		root.Left = insert(root.Left, v)
		if root.GetBF() == 2 {
			if v < root.Left.Value {
				// LL condition
				root = RRotate(root)
			} else {
				// LR condition
				root = LRRotate(root)
			}
		}
	} else if v > root.Value {
		root.Right = insert(root.Right, v)
		if root.GetBF() == 2 {
			if v > root.Left.Value {
				// RR condition
				root = LRotate(root)
			} else {
				// RL condition
				root = RLRotate(root)
			}
		}
	}
	// recaclulate height
	root.height = max(getHeight(root.Left), getHeight(root.Right)) + 1
	return root
}

// Delete a tree node and keep balance
func (t *AVLTree) Delete(v int) {
	t.root = delete(t.root, v)
}

func delete(root *Node, v int) *Node {
	if root == nil {
		return nil
	}
	if root.Value == v {
		// delete first
		if root.Left != nil && root.Right != nil {
			if getHeight(root.Left) > getHeight(root.Right) {
				tmpr := FindMax(root.Left)
				root.Value = tmpr.Value
				root.Left = delete(root.Left, tmpr.Value)
			} else {
				tmpr := FindMin(root.Right)
				root.Value = tmpr.Value
				root.Right = delete(root.Right, tmpr.Value)
			}
		} else {
			if root.Left != nil {
				root = root.Left
			} else if root.Right != nil {
				root = root.Right
			} else {
				root = nil
			}
		}
		return root
	} else if root.Value < v {
		// and then rotate
		root.Right = delete(root.Right, v)
		if root.GetBF() == 2 {
			if v < root.Left.Value {
				// LL condition
				root = RRotate(root)
			} else if v > root.Left.Value {
				// LR condition
				root = LRRotate(root)
			}
		}
	} else if root.Value > v {
		root.Left = delete(root.Left, v)
		if root.GetBF() == 2 {
			if v > root.Right.Value {
				// RR condition
				root = LRotate(root)
			} else if v < root.Right.Value {
				// RL condition
				root = RLRotate(root)
			}
		}
	}
	return root
}

// FindMax returns the biggest node on the tree
func FindMax(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.Right != nil {
		return FindMax(n.Right)
	}
	return n
}

// FindMin returns the smallest node on the tree
func FindMin(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.Left != nil {
		return FindMin(n.Left)
	}
	return n
}
