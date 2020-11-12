/*
Package fenwicktree implements a  binary indexed tree:
	A **Fenwick tree** or **binary indexed tree** is a data
	structure that can efficiently update elements and
	calculate prefix sums in a table of numbers.

	When compared with a flat array of numbers, the Fenwick tree achieves a
	much better balance between two operations: element update and prefix sum
	calculation. In a flat array of `n` numbers, you can either store the elements,
	or the prefix sums. In the first case, computing prefix sums requires linear
	time; in the second case, updating the array elements requires linear time
	(in both cases, the other operation can be performed in constant time).
	Fenwick trees allow both operations to be performed in `O(log n)` time.
	This is achieved by representing the numbers as a tree, where the value of
	each node is the sum of the numbers in that subtree. The tree structure allows
	operations to be performed using only `O(log n)` node accesses.
WikiPage:
	* https://en.wikipedia.org/wiki/Fenwick_tree
*/
package fenwicktree

// FenWickTree is a binary indexed tree data structure
type FenWickTree struct {
	size int
	tree []int
}

// New construct a fenwick tree with an array size s
// however, the actuall array size is size+1, because
// index is 1-based.
func New(s int) *FenWickTree {
	return &FenWickTree{
		size: s + 1,
		tree: make([]int, s+1),
	}
}

// Increase update ele at i to v+num[i].
func (t *FenWickTree) Increase(i, v int) {
	if i < 1 || i > t.size {
		return
	}
	for ; i < t.size; i += lowbit(i) {
		t.tree[i] += v
	}
}

// return 2^k, k = last bit 1`offset
func lowbit(i int) int {
	return i & (-i)
}

// Query sum from index 1 to position.
func (t *FenWickTree) Query(i int) int {
	if i < 1 || i > t.size {
		panic("out of bounds")
	}
	sum := 0
	for ; i > 0; i -= lowbit(i) {
		sum += t.tree[i]
	}
	return sum
}

// Range sum from index l ro index r.
func (t *FenWickTree) Range(l, r int) int {
	if l < r {
		panic("invalid bounds")
	}
	if l == 1 {
		return t.Query(r)
	}
	return t.Query(r) - t.Query(l)
}
