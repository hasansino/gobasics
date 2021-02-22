package searching

// https://en.wikipedia.org/wiki/Binary_search_algorithm

func BinarySearch(data []int, s int) bool {
	l := 0
	r := len(data) - 1
	for l <= r {
		m := (l + r) / 2
		switch {
		case data[m] < s:
			l = m + 1
		case data[m] > s:
			r = m - 1
		default:
			return true
		}
	}
	return false
}
