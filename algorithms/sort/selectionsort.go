package sort

// SelectSort is the worst
func SelectSort(data Sortable) {
	for i := 0; i < data.Len(); i++ {
		min := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(min, i)
	}
}

/*
Complexity of select sort：
	* Best: 	O(n^2)
	* Average: 	O(n^2)
	* Worst: 	O(n^2)
	* Memory: 	O(1)
	* Stable: 	No
	* Wiki: 	https://en.wikipedia.org/wiki/Selection_sort
Shortcome from wiki:
	Selection sort is a sorting algorithm, specifically an
	in-place comparison sort. It has O(n2) time complexity,
	making it inefficient on large lists, and generally
	performs worse than the similar insertion sort.
	Selection sort is noted for its simplicity, and it has
	performance advantages over more complicated algorithms
	in certain situations, particularly where auxiliary
	memory is limited.
*/
