package sorting

import "sort"

// https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort

func OddEvenSort(data sort.Interface) {
	var sorted bool
	for !sorted {
		sorted = true
		// even indexes
		for i := 0; i < data.Len()-1; i += 2 {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
				sorted = false
			}
		}
		// odd indexes
		for i := 1; i < data.Len()-1; i += 2 {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
				sorted = false
			}
		}
	}
}
