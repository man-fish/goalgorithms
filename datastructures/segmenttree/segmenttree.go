/*
Package segmenttree implements a segment tree:
	In computer science, a **segment tree** also known as a statistic tree
	is a tree data structure used for storing information about intervals,
	or segments. It allows querying which of the stored segments contain
	a given point. It is, in principle, a static structure; that is,
	it's a structure that cannot be modified once it's built. A similar
	data structure is the interval tree.

	A segment tree is a binary tree. The root of the tree represents the
	whole array. The two children of the root represent the
	first and second halves of the array. Similarly, the
	children of each node corresponds to the two halves of
	the array corresponding to the node.

	We build the tree bottom up, with the value of each node
	being the "minimum" (or any other function) of its children's values. This will
	take `O(n log n)` time. The number
	of operations done is the height of the tree, which
	is `O(log n)`. To do range queries, each node splits the
	query into two parts, one sub-query for each child.
	If a query contains the whole subarray of a node, we
	can use the precomputed value at the node. Using this
	optimisation, we can prove that only `O(log n)` minimum
	operations are done.
WikiPage:
	* https://en.wikipedia.org/wiki/Segment_tree
*/
package segmenttree

// SegmentTree represents the segment tree datastructure
type SegmentTree struct {
	// segment arr tree
	tree []int
	// datas
	arr []int
}

// New is a constructor
// 4*n is tree max length
func New(arr []int) *SegmentTree {
	n := len(arr)
	t := &SegmentTree{
		tree: make([]int, 4*n),
		arr:  make([]int, n),
	}
	t.arr = arr
	t.build(0, 0, n-1)
	return t
}

// n represents cur build node, [s, e] represent [start, end] of segment in node
func (t *SegmentTree) build(n int, s, e int) {
	if s == e {
		t.tree[n] = t.arr[s]
	} else {
		m := (s + e) / 2
		ln := 2*n + 1
		rn := 2*n + 2
		t.build(ln, s, m)
		t.build(rn, m+1, e)
		t.tree[n] = t.tree[ln] + t.tree[rn]
	}
}

// Update update a single in segment tree
// n represents node, [s, e] represent [start, end] of segment in node,
// i represents idx in arr, v represents new node value
func (t *SegmentTree) Update(n, s, e, i, v int) {
	if s == e {
		t.arr[i] = v
		t.tree[n] = v
	} else {
		m := (s + e) / 2
		ln := 2*n + 1
		rn := 2*n + 2
		if i >= s && i <= m {
			t.Update(ln, s, m, i, v)
		} else {
			t.Update(rn, m+1, e, i, v)
		}
		t.tree[n] = t.tree[ln] + t.tree[rn]
	}
}

// Sum returns certain segment sum in segment tree
// n represents node, [s, e] represent [start, end] of segment in node,
// [l, r] represent [right, left] segment in arr to sum
func (t *SegmentTree) Sum(n, s, e, l, r int) int {
	if l > e || r < s {
		return 0
	} else if l <= s && e <= r {
		return t.tree[n]
	} else if s == e {
		return t.tree[n]
	} else {
		m := (s + e) / 2
		ln := 2*n + 1
		rn := 2*n + 2
		sl := t.Sum(ln, s, m, l, r)
		sr := t.Sum(rn, m+1, e, l, r)
		return sl + sr
	}
}
