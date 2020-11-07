package sort

// MergeSort is a O(nlog(n)) stable sorting algorithm
func MergeSort(data []int) {
	mergeSort(data, make([]int, len(data)), 0, len(data)-1)
}

func mergeSort(data, temp []int, start, end int) {
	if start < end {
		mid := int(uint(start+end) >> 1)
		mergeSort(data, temp, start, mid)
		mergeSort(data, temp, mid+1, end)
		merge(data, temp, start, mid, end)
	}
}

func merge(data, temp []int, start, mid, end int) {
	i := start
	j := mid + 1
	t := 0
	for i <= mid && j <= end {
		if data[i] < data[j] {
			temp[t] = data[i]
			i++
			t++
		} else {
			temp[t] = data[j]
			j++
			t++
		}
	}

	for i <= mid {
		temp[t] = data[i]
		i++
		t++
	}

	for j <= end {
		temp[t] = data[j]
		j++
		t++
	}

	t = 0
	for start <= end {
		data[start] = temp[t]
		t++
		start++
	}
}

/*
Complexity of merge sort：
	* Best: 	O(nlog(n))
	* Average: 	O(nlog(n))
	* Worst: 	O(nlog(n))
	* Memory: 	O(n)
	* Stable: 	Yes
	* Wiki: 	https://en.wikipedia.org/wiki/Merge_sort
Shortcome from wiki:
	In computer science, merge sort (also commonly spelled
	mergesort) is an efficient, general-purpose,
	comparison-based sorting algorithm. Most implementations
	produce a stable sort, which means that the implementation
	preserves the input order of equal elements in the sorted
	output. Mergesort is a divide and conquer algorithm that
	was invented by John von Neumann in 1945.

	An example of merge sort. First divide the list into
	the smallest unit (1 element), then compare each
	element with the adjacent list to sort and merge the
	two adjacent lists. Finally all the elements are sorted
	and merged.
*/
