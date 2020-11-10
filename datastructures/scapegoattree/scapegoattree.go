package scapegoattree

const alpha = 0.75

type Node struct {
	Left, Right *Node
	Val         int
	// size represent efftive nodes under this node
	// cover represent total nodes nums
	size, cover int
	// exist marks whether the node is lazy deleted
	exist bool
}
