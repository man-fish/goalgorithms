package sort

// BubbleSort is a O(n^2) stable sorting algorithm
func BubbleSort(data Sortable) {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len()-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

/*
Complexity of bubble sort：
	* Best: 	O(n)
	* Average: 	O(n^2)
	* Worst: 	O(n^2)
	* Memory: 	O(1)
	* Stable: 	Yes
	* Wiki: 	https://en.wikipedia.org/wiki/Bubble_sort
Shortcome from wiki:
	Bubble sort, sometimes referred to as sinking sort, is a
	simple sorting algorithm that repeatedly steps through
	the list to be sorted, compares each pair of adjacent
	items and swaps them if they are in the wrong order
	(ascending or descending arrangement). The pass through
	the list is repeated until no swaps are needed, which
	indicates that the list is sorted.
*/
