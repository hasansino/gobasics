package searching

// https://en.wikipedia.org/wiki/Binary_search_algorithm

func BinarySearch(data, dataSorted []int, s int) bool {
	l := 0
	r := len(dataSorted)
	for l <= r {
		m := (l + r) / 2
		switch {
		case dataSorted[m] < s:
			l = m + 1
		case dataSorted[m] > s:
			r = m - 1
		default:
			return true
		}
	}
	return false
}
