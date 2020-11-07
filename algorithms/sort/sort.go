package sort

import "math/rand"

// Sortable represent the elements which can be sorted.
type Sortable interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Equal reports whether the element is equal
	Equal(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

func (p IntSlice) Len() int            { return len(p) }
func (p IntSlice) Less(i, j int) bool  { return p[i] < p[j] }
func (p IntSlice) Equal(i, j int) bool { return p[i] == p[j] }
func (p IntSlice) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }

var _ Sortable = (*IntSlice)(nil)

// RandomArray generate a random arr
func RandomArray(length, max int) IntSlice {
	arr := make([]int, length)
	for idx := range arr {
		arr[idx] = rand.Intn(max)
	}
	return IntSlice(arr)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
