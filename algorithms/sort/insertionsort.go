package sort

// InsertionSort is a O(n^2) stable sorting algorithm
func InsertionSort(data Sortable) {
	for i := 1; i < data.Len(); i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

/*
Complexity of insertion sort：
	* Best: 	O(n)
	* Average: 	O(n^2)
	* Worst: 	O(n^2)
	* Memory: 	O(1)
	* Stable: 	Yes
	* Wiki: 	https://en.wikipedia.org/wiki/Insertion_sort
Shortcome from wiki:
	Insertion sort is a simple sorting algorithm that builds
	the final sorted array (or list) one item at a time.
	It is much less efficient on large lists than more
	advanced algorithms such as quicksort, heapsort, or merge
	sort.
*/
