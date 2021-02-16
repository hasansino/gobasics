package sorting

import "sort"

// https://en.wikipedia.org/wiki/Quicksort

func QuickSort(data sort.Interface) {
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data sort.Interface, lo, hi int) {
	if lo < hi {
		p := quickSortPartition(data, lo, hi)
		quickSort(data, lo, p-1)
		quickSort(data, p+1, hi)
	}
}

func quickSortPartition(data sort.Interface, lo, hi int) int {
	pivot := quickSortPivot(data, lo, hi)
	i := lo // partition point
	for j := lo; j < hi; j++ {
		if data.Less(j, pivot) {
			data.Swap(i, j)
			i++
		}
	}
	data.Swap(i, hi)
	return i
}

// https://en.wikipedia.org/wiki/Quicksort#Choice_of_pivot
func quickSortPivot(data sort.Interface, lo, hi int) int {
	return hi
}
