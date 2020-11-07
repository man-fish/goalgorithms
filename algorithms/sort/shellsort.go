package sort

// ShellSort is a unstable sorting algorithm
func ShellSort(data Sortable) {
	for dalta := data.Len() / 2; dalta > 0; dalta = dalta >> 1 {
		for i := dalta; i < data.Len(); i++ {
			for j := i; j-dalta >= 0 && data.Less(j, j-dalta); j -= dalta {
				data.Swap(j, j-dalta)
			}
		}
	}
}

/*
Complexity of shell sort：
	* Best: 	O(nlog(n))
	* Average: 	depends on gap sequence
	* Worst: 	O(nlog(n)^2)
	* Memory: 	O(1)
	* Stable: 	No
	* Wiki: 	https://en.wikipedia.org/wiki/Shellsort
Shortcome from wiki:
	Shellsort, also known as Shell sort or Shell's method,
	is an in-place comparison sort. It can be seen as either a
	generalization of sorting by exchange (bubble sort) or sorting
	by insertion (insertion sort). The method starts by sorting
	pairs of elements far apart from each other, then progressively
	reducing the gap between elements to be compared. Starting
	with far apart elements, it can move some out-of-place
	elements into position faster than a simple nearest neighbor
	exchange
*/
