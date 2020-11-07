package sort

import (
	"math"
	"strconv"
)

// RadixSort is a O(n*k) stable sorting algorithm
func RadixSort(intArr []int) []int {
	tmp := make([]int, len(intArr))
	copy(tmp, intArr)
	digits := MaxDigits(tmp)
	bucket := make([]int, digits)

	for index := range bucket {
		zeros := int(math.Pow(float64(10), float64(index)))
		count := [10]int{}
		intArr = rearange(zeros, intArr, &count)
	}
	return intArr
}

// MaxDigits returns the biggest digit among nums
func MaxDigits(intArr []int) int {
	for i := 1; i < len(intArr); i++ {
		if intArr[i-1] > intArr[i] {
			intArr[i-1], intArr[i] = intArr[i], intArr[i-1]
		}
	}
	maxNum := intArr[len(intArr)-1]
	return len(strconv.Itoa(maxNum))
}

func rearange(zores int, intArr []int, count *[10]int) []int {
	for _, value := range intArr {
		(*count)[(value/zores)%10]++
	}
	(*count)[0]--
	for i := 1; i < len(count); i++ {
		(*count)[i] = (*count)[i-1] + (*count)[i]
	}

	aux := make([]int, len(intArr))
	for i := len(intArr) - 1; i >= 0; i-- {
		dnum := (intArr[i] / zores) % 10
		pos := (*count)[dnum]
		aux[pos] = intArr[i]
		(*count)[dnum]--
	}
	return aux
}

/*
Complexity of select sort：
	* Best: 	O(n*k)
	* Average: 	O(n*k)
	* Worst: 	O(n*k)
	* Memory: 	O(n+k)
	* Stable: 	Yes
	* Wiki: 	https://en.wikipedia.org/wiki/Radix_sort
Shortcome from wiki:
	In computer science, radix sort is a non-comparative integer sorting
	algorithm that sorts data with integer keys by grouping keys by the individual
	digits which share the same significant position and value. A positional notation
	is required, but because integers can represent strings of characters
	(e.g., names or dates) and specially formatted floating point numbers, radix
	sort is not limited to integers.

	*Where does the name come from?*
		In mathematical numeral systems, the *radix* or base is the number of unique digits,
		including the digit zero, used to represent numbers in a positional numeral system.
		For example, a binary system (using numbers 0 and 1) has a radix of 2 and a decimal
		system (using numbers 0 to 9) has a radix of 10.
*/
