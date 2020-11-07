package sort

import "github.com/man-fish/goalgorithms/datastructures/heap"

// HeapSort is a O(nlog(n)) stable sorting algorithm
func HeapSort(data []int) []int {
	return heap.Sort(data)
}

/*
Complexity of merge sort：
	* Best: 	O(nlog(n))
	* Average: 	O(nlog(n))
	* Worst: 	O(nlog(n))
	* Memory: 	O(1)
	* Stable: 	No
	* Wiki: 	https://en.wikipedia.org/wiki/Heapsort
Shortcome from wiki:
	Heapsort is a comparison-based sorting algorithm.
	Heapsort can be thought of as an improved selection
	sort: like that algorithm, it divides its input into
	a sorted and an unsorted region, and it iteratively
	shrinks the unsorted region by extracting the largest
	element and moving that to the sorted region. The
	improvement consists of the use of a heap data structure
	rather than a linear-time search to find the maximum.
*/
