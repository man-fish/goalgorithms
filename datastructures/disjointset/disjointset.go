/*
Package disjointset implements a doubly disjoint set:
	**Disjoint-set** data structure (also called a union–find data structure or merge–find set) is a data
	structure that tracks a set of elements partitioned into a number of disjoint (non-overlapping) subsets.
	It provides near-constant-time operations (bounded by the inverse Ackermann function) to *add new sets*,
	to *merge existing sets*, and to *determine whether elements are in the same set*.
	In addition to many other uses (see the Applications section), disjoint-sets play a key role in Kruskal's
	algorithm for finding the minimum spanning tree of a graph.
WikiPage:
	* https://en.wikipedia.org/wiki/Disjoint-set_data_structure
*/
package disjointset

import (
	"errors"
	"fmt"
)

// DisjointSet represent a disjointset struct
type DisjointSet struct {
	rank    []int
	parents []int
}

// New is a constructor
func New(size int) *DisjointSet {
	djset := &DisjointSet{
		parents: make([]int, size),
		rank:    make([]int, size),
	}
	for i := 0; i < size; i++ {
		djset.parents[i] = -1
		djset.rank[i] = 0
	}
	return djset
}

// Size return vertices num of djset
func (s *DisjointSet) Size() int {
	return len(s.parents)
}

// Union check and merge two existing sets
// if set a and set b create a cycle, union will fail and return 0
// else two set will be merge to be a smaller tree
func (s *DisjointSet) Union(a, b int) (int, error) {
	if a < 0 || a >= s.Size() || b < 0 || b >= s.Size() {
		return 0, errors.New("index out of bounds")
	}
	ra := s.findParent(a)
	rb := s.findParent(b)
	fmt.Println(ra, rb)
	if ra == rb {
		return 0, nil
	}
	// union root
	if s.rank[ra] > s.rank[rb] {
		s.parents[rb] = ra
	} else if s.rank[ra] < s.rank[rb] {
		s.parents[ra] = rb
	} else {
		s.parents[rb] = ra
		s.rank[ra]++
	}
	return 1, nil
}

func (s *DisjointSet) findParent(i int) int {
	p := i
	for s.parents[p] != -1 {
		p = s.parents[p]
	}
	return p
}
