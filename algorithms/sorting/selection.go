package sorting

import "sort"

// https://en.wikipedia.org/wiki/Selection_sort

func SelectionSort(data sort.Interface) {
	for i := 0; i < data.Len()-1; i++ {
		min := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		if min != i {
			data.Swap(min, i)
		}
	}
}
