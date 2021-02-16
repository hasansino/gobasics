package sorting

import "sort"

// https://en.wikipedia.org/wiki/Gnome_sort

func GnomeSort(data sort.Interface) {
	for i := 0; i < data.Len(); {
		if i == 0 {
			i++
			continue
		}
		if data.Less(i, i-1) {
			data.Swap(i, i-1)
			i--
		} else {
			i++
		}
	}
}
