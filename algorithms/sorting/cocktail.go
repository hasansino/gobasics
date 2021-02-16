package sorting

import "sort"

// https://en.wikipedia.org/wiki/Cocktail_shaker_sort

func CocktailSort(data sort.Interface) {
	var pass int
	for {
		pass++
		dryRun := true
		if pass%2 == 0 { // even passes go from start to end of slice
			for i := 0; i < data.Len()-1; i++ {
				if data.Less(i+1, i) {
					data.Swap(i, i+1)
					dryRun = false
				}
			}
		} else { // odd passes go from end to start of slice
			for i := data.Len() - 1; i > 0; i-- {
				if data.Less(i, i-1) {
					data.Swap(i-1, i)
					dryRun = false
				}
			}
		}
		if dryRun {
			break
		}
	}
}
