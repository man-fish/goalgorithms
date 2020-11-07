package sort_test

import (
	"sort"
	"testing"

	isort "github.com/man-fish/goalgorithms/algorithms/sort"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var positiveInts = []int{74, 59, 238, 9845, 959, 905, 0, 0, 42, 7586, 7586}

func TestBubbleSort(t *testing.T) {
	data := ints
	a := isort.IntSlice(data[0:])
	isort.BubbleSort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf(" goted %v", data)
	}
}

func TestCountingSort(t *testing.T) {
	sorted := isort.CountingSort(positiveInts)
	sortedSlice := isort.IntSlice(sorted)
	if !sort.IsSorted(sortedSlice) {
		t.Errorf("sorted %v", positiveInts)
		t.Errorf(" goted %v", sorted)
	}
}

func TestHeapSort(t *testing.T) {
	sorted := isort.HeapSort(ints[0:])
	sortedSlice := isort.IntSlice(sorted)
	if !sort.IsSorted(sortedSlice) {
		t.Errorf("sorted %v", positiveInts)
		t.Errorf(" goted %v", sorted)
	}
}

func TestRadixSort(t *testing.T) {
	sorted := isort.RadixSort(positiveInts)
	sortedSlice := isort.IntSlice(sorted)
	if !sort.IsSorted(sortedSlice) {
		t.Errorf("sorted %v", positiveInts)
		t.Errorf(" goted %v", sorted)
	}
}

func TestInsertionSort(t *testing.T) {
	data := ints
	a := isort.IntSlice(data[0:])
	isort.InsertionSort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf(" goted %v", data)
	}
}

func TestMergeSort(t *testing.T) {
	data := ints
	a := isort.IntSlice(data[0:])
	isort.MergeSort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf(" goted %v", data)
	}
}

func TestQuickSort(t *testing.T) {
	data := ints
	a := isort.IntSlice(data[0:])
	isort.QuickSort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf(" goted %v", data)
	}
}

func TestSelectSort(t *testing.T) {
	data := ints
	a := isort.IntSlice(data[0:])
	isort.SelectSort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf(" goted %v", data)
	}
}

func TestShellSort(t *testing.T) {
	data := ints
	a := isort.IntSlice(data[0:])
	isort.ShellSort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf(" goted %v", data)
	}
}
