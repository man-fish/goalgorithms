package sort

// QuickSort is a O(nlog(n)) stable sorting algorithm
func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, lo, hi int) {
	if len(data) == 0 {
		return
	}
	mlo, mhi := lo, hi
	pivot := data[int(uint(mlo+mhi)>>1)]

	for mlo < mhi {
		for data[mlo] < pivot {
			mlo++
		}
		for pivot < data[mhi] {
			mhi--
		}
		if mhi <= mlo {
			break
		}
		data[mlo], data[mhi] = data[mhi], data[mlo]
		if data[mlo] == pivot {
			mhi--
		}
		if data[mhi] == pivot {
			mlo++
		}
	}

	if mlo == mhi {
		mlo++
		mhi--
	}
	if mlo < hi {
		quickSort(data, mlo, hi)
	}
	if mhi > lo {
		quickSort(data, lo, mhi)
	}
}

/*
Complexity of quick sort：
	* Best: 	O(nlog(n))
	* Average: 	O(nlog(n))
	* Worst: 	O(nlog(n))
	* Memory: 	O(n^2)
	* Stable: 	Yes
	* Wiki: 	https://en.wikipedia.org/wiki/Quicksort
Shortcome from wiki:
	Quicksort is a divide and conquer algorithm.
	Quicksort first divides a large array into two smaller
	sub-arrays: the low elements and the high elements.
	Quicksort can then recursively sort the sub-arrays
*/
