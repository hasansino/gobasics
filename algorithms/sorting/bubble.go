package sorting

import "sort"

// https://en.wikipedia.org/wiki/Bubble_sort

func BubbleSort(data sort.Interface) {
	for {
		dryRun := true
		for i := 0; i < data.Len()-1; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
				dryRun = false
			}
		}
		if dryRun {
			break
		}
	}
}
