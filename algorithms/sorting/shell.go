package sorting

import "sort"

// https://en.wikipedia.org/wiki/Shellsort

func ShellSort(data sort.Interface) {
	for gap := data.Len() / 2; gap > 0; gap /= 2 {
		for i := gap; i < data.Len(); i++ {
			for j := i; j >= gap && data.Less(j, j-gap); j -= gap {
				data.Swap(j, j-gap)
			}
		}
	}
}
