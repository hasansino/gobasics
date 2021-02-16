package sorting

import "sort"

// https://en.wikipedia.org/wiki/Insertion_sort

func InsertionSort(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		j := i
		for j > 0 { // first iteration is skipped
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
			j--
		}
	}
}
