package sorting

import "sort"

// https://en.wikipedia.org/wiki/Comb_sort

func CombSort(data sort.Interface) {
	const shrinkFactor = 1.3
	gap := data.Len()
	for sorted := false; !sorted; {
		// conversion to int will cut decimal part
		gap = int(float64(gap) / shrinkFactor)
		if gap <= 1 {
			gap = 1
			sorted = true
		}
		for i := 0; i+gap < data.Len(); i++ {
			if data.Less(i+gap, i) {
				data.Swap(i, i+gap)
				sorted = false
			}
		}
	}
}
