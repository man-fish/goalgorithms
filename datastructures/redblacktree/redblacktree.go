/*
Package redblacktree implements a red-black tree:
	A **red–black tree** is a kind of self-balancing binary search
	tree in computer science. Each node of the binary tree has
	an extra bit, and that bit is often interpreted as the
	color (red or black) of the node. These color bits are used
	to ensure the tree remains approximately balanced during
	insertions and deletions.

	Balance is preserved by painting each node of the tree with
	one of two colors in a way that satisfies certain properties,
	which collectively constrain how unbalanced the tree can
	become in the worst case. When the tree is modified, the
	new tree is subsequently rearranged and repainted to
	restore the coloring properties. The properties are
	designed in such a way that this rearranging and recoloring
	can be performed efficiently.

	The balancing of the tree is not perfect, but it is good
	enough to allow it to guarantee searching in `O(log n)` time,
	where `n` is the total number of elements in the tree.
	The insertion and deletion operations, along with the tree
	rearrangement and recoloring, are also performed
	in `O(log n)` time.
WikiPage:
	* https://en.wikipedia.org/wiki/Red%E2%80%93black_tree
*/
package redblacktree

const (
	red int8 = iota
	black
)

// Node is an element of red-black-tree
type Node struct {
	value  int
	color  int8
	left   *Node
	right  *Node
	parent *Node
}

func (n *Node) setLeft(l *Node) {
	if n.left != nil {
		n.left.parent = nil
	}
	n.left = l
	if n.left != nil {
		n.left.parent = n
	}
}

func (n *Node) setRight(r *Node) {
	if n.left != nil {
		n.right.parent = nil
	}
	n.right = r
	if n.right != nil {
		n.right.parent = n
	}
}

func (n *Node) uncle() *Node {
	if n.parent != nil && n.parent.parent != nil {
		grandpa := n.parent.parent
		if n.parent == grandpa.left {
			return grandpa.right
		}
		return grandpa.left
	}
	return nil
}

func (n *Node) bro() *Node {
	if n.parent != nil {
		if n == n.parent.left {
			return n.parent.right
		}
		return n.parent.left
	}
	return nil
}

func (n *Node) paint(color int8) {
	n.color = color
}

func (n *Node) isRed() bool {
	return n.color == red
}

func (n *Node) isBlack() bool {
	return n.color == black
}

func (n *Node) isPainted() bool {
	return n.isRed() || n.isBlack()
}

func (n *Node) exchangeColor(o *Node) {
	if o != nil {
		n.color, o.color = o.color, n.color
	}
}

func lRotate(n *Node) *Node {
	// Memorize grandParentNode's right node.
	parent := n.right
	// Move child's left subtree to grandParentNode's right subtree.
	n.setRight(parent.left)
	// Make grandParentNode to be left child of parentNode.
	parent.setLeft(parent)
	return parent
}

func rRotate(n *Node) *Node {
	// Memorize grandParentNode's left node.
	parent := n.left
	// Move child's right subtree to grandParentNode's left subtree.
	n.setLeft(parent.right)
	// // grandpa.left = parent.right
	// // parent.right.parent = grandpa
	// Make grandParentNode to be right child of parentNode.
	parent.setRight(n)
	return parent
}

// RedBlackTree represents a red-black-tree
type RedBlackTree struct {
	root *Node
}

// Search returns whether is in the tree
func (t *RedBlackTree) Search(v int) bool {
	return t.search(t.root, v)
}

func (t *RedBlackTree) search(root *Node, v int) bool {
	if root == nil {
		return false
	}

	if root.value == v {
		return true
	}

	if root.value > v {
		return t.search(root.left, v)
	}
	return t.search(root.right, v)
}

// Insert insert a node and balance the tree, at last return the inserted node
func (t *RedBlackTree) Insert(v int) *Node {
	if t.root == nil {
		t.root = &Node{value: v, color: black}
		return t.root
	}
	n := t.insert(t.root, v)
	if n != nil {
		t.balance(n)
	}
	return n
}

func (t *RedBlackTree) insert(r *Node, v int) *Node {
	if r.value == v {
		return nil
	}
	if v > r.value {
		if r.right == nil {
			// color is default 0 - red
			n := &Node{value: v}
			r.setRight(&Node{value: v})
			return n
		}
		return t.insert(r.right, v)
	} else {
		if r.left == nil {
			// color is default 0 - red
			n := &Node{value: v}
			r.setLeft(n)
			return n
		}
		return t.insert(r.right, v)
	}
}

func (t *RedBlackTree) balance(n *Node) {
	if n == nil || t.root == n {
		// If it is a root node or nil then nothing to balance here.
		return
	}
	if n.parent.isBlack() {
		// If the parent is black then done. Nothing to balance here.
		// only two red nodes linked wil break the rule
		return
	}
	// now n must have a red parent but may not have grandpa or uncle
	// and n itself is default red
	parent, uncle := n.parent, n.uncle()
	grandpa := n.parent.parent
	if uncle != nil && uncle.isRed() {
		// If node has red uncle then we need to do RECOLORING.

		// Recolor parent and uncle to black.
		parent.paint(black)
		uncle.paint(black)
		if grandpa != t.root {
			// Recolor grand-parent to red if it is not root.
			grandpa.paint(red)
		} else {
			// If grand-parent is black root don't do anything.
			// Since root already has two black sibling that we've just recolored.
			// balance with black(grandpa) black(parent) red(itslef)
			return
		}
		// Now do further checking for recolored grand-parent.
		t.balance(grandpa)
	} else if uncle == nil || uncle.isBlack() {
		// now node is red then:
		// If node uncle is black or absent then we need to do ROTATIONS.
		if grandpa != nil {
			// Grand parent that we will receive after rotations.
			var npa *Node
			if grandpa.left == parent {
				// on left
				if n.parent.left == n {
					// on left left
					npa = t.llRotate(grandpa)
				} else {
					// on left right
					npa = t.lrRotate(grandpa)
				}
			} else {
				// on right
				if n.parent.left == n {
					// on right left
					npa = t.rlRotate(grandpa)
				} else {
					// on right right
					npa = t.rrRotate(grandpa)
				}
			}

			if npa != nil && npa.parent == nil {
				// Set newGrandParent as a root if it doesn't have parent.
				t.root = npa
				// Recolor root into black.
				t.root.paint(black)
			}
			// Check if new grand parent don't violate red-black-tree rules.
			t.balance(npa)
		}
	}
}

func (t *RedBlackTree) llRotate(grandpa *Node) *Node {
	// Memorize the parent of grand-parent node.
	grandpapa := grandpa.parent
	// Check what type of sibling is our grandParentNode is (left or right).
	grandpaAtLeft := grandpapa != nil && grandpa == grandpapa.left
	parent := rRotate(grandpa)
	if grandpapa != nil {
		if grandpaAtLeft {
			grandpapa.setLeft(parent)
		} else {
			grandpapa.setRight(parent)
		}
	} else {
		// Make parent node a root
		parent.parent = nil
	}
	// Swap colors of granParent and parent nodes.
	parent.exchangeColor(grandpa)
	// Return new root node.
	return parent
}

func (t *RedBlackTree) lrRotate(grandpa *Node) *Node {
	// Memorize left and left-right nodes.
	parent := grandpa.left
	rSon := parent.right
	// left child subtree will be re-assigned to parent's right sub-tree.
	parent.setRight(rSon.left)
	// // parent.right = rSon.left
	// // rSon.left.parent = parent
	// Make parentNode to be a left child of childNode node.
	rSon.setLeft(parent)
	// //rSon.left = parent
	// //parent.parent = rSon
	// Put left-right node in place of left node.
	grandpa.setLeft(rSon)
	// // grandpa.left = rSon
	// // rSon.parent = grandpa
	// now it is still not balance we need to do left-left rotation.
	return t.llRotate(grandpa)
}

func (t *RedBlackTree) rrRotate(grandpa *Node) *Node {
	// Memorize the parent of grand-parent node.
	grandpapa := grandpa.parent
	// Check what type of sibling is our grandParentNode is (left or right).
	grandpaAtLeft := grandpapa != nil && grandpa == grandpapa.left
	parent := lRotate(grandpa)
	if grandpapa != nil {
		if grandpaAtLeft {
			grandpapa.setLeft(parent)
		} else {
			grandpapa.setRight(parent)
		}
	} else {
		// Make parent node a root
		parent.parent = nil
	}
	// Swap colors of granParent and parent nodes.
	parent.exchangeColor(grandpa)
	// Return new root node.
	return parent
}

func (t *RedBlackTree) rlRotate(grandpa *Node) *Node {
	// Memorize right and right-left nodes.
	parent := grandpa.right
	lSon := parent.left
	// right child subtree  will be re-assigned to  parent's left sub-tree.
	parent.setLeft(lSon.right)
	// // parent.left = rSon.right
	// // rSon.right.parent = parent
	// Make parentNode to be a right child of childNode.
	lSon.setRight(parent)
	// //rSon.right = parent
	// //parent.parent = lSOn
	// Put childNode node in place of parentNode.
	grandpa.setRight(lSon)
	// // grandpa.right = lSon
	// // lSon.parent = grandpa
	// now it is still not balance we need to do right-right rotation.
	return t.rrRotate(grandpa)
}

// Delete an element from the tree and keep the tree balance
func (t *RedBlackTree) Delete(v int) {
	t.delete(t.root, v)
}

// delete return the new element on the deleted position
func (t *RedBlackTree) delete(root *Node, v int) *Node {
	if root == nil {
		return nil
	}

	if root.value == v {
		if root.left != nil && root.right != nil {
			cur := root
			next := root.left

			for next.right != nil {
				// find the biggest node before deleted node
				cur = next
				next = next.right
			}

			// exchange value and color
			root.value = next.value
			root.color = next.color

			if cur == root {
				cur.setLeft(next.left)
			} else {
				cur.setRight(next.left)
			}
		} else {
			if root.left != nil {
				root = root.left
			} else {
				root = root.right
			}
		}
		if root.color == black {
			t.eliminate(root)
		}
		return root
	} else if root.value > v {
		root.setLeft(t.delete(root.left, v))
	} else {
		root.setRight(t.delete(root.right, v))
	}
	return root
}

func (t *RedBlackTree) eliminate(n *Node) {
	var parent, bro *Node
	for (n == nil || n != t.root) && n.isBlack() {
		parent = n.parent
		bro = n.bro()
		if parent.left == n {
			if bro.isRed() {
				// case1: n is black+black, bro is red
				bro.paint(black)
				parent.paint(red)
				lRotate(parent)
				// set x`s new bro
				bro = parent.right
			}

			if (bro.left == nil || bro.left.isBlack()) && (bro.right == nil || bro.right.isBlack()) {
				// case2: n is black+black, bro is black, bro`left and right are black or nil
				bro.color = red
				n = parent
				parent = n.parent
			} else {
				if bro.right == nil || bro.right.isBlack() {
					// case3: n is black+black, bro is black, bro`right is black and left is red
					bro.left.paint(black)
					bro.paint(red)
					rRotate(bro)
					bro = parent.right
				}
				// case4: n is black+black, bro is black, bro`left is any color and right is red
				bro.paint(parent.color)
				parent.paint(black)
				bro.right.paint(black)
				lRotate(parent)
				n = t.root
				break
			}
		} else {
			if bro.isRed() {
				// case1: n is black+black, bro is red
				bro.paint(black)
				parent.paint(red)
				rRotate(parent)
				// set x`s new bro
				bro = parent.left
			}

			if (bro.left == nil || bro.left.isBlack()) && (bro.right == nil || bro.right.isBlack()) {
				// case2: n is black+black, bro is black, bro`left and right are black or nil
				bro.paint(red)
				n = parent
				parent = n.parent
			} else {
				if bro.right == nil || bro.right.isBlack() {
					// case3: n is black+black, bro is black, bro`right is black and left is red
					bro.right.paint(black)
					bro.paint(red)
					lRotate(bro)
					bro = parent.left
				}
				// case4: n is black+black, bro is black, bro`left is any color and right is red
				bro.paint(parent.color)
				parent.paint(black)
				bro.left.paint(black)
				rRotate(parent)
				n = t.root
				break
			}
		}
	}
	if n != nil {
		n.paint(red)
	}
}
