package searchnum

// https://en.wikipedia.org/wiki/Exponential_search

func ExponentialSearch(data []int, s int) bool {
	if len(data) == 0 {
		return false
	}
	bound := 1
	for bound < len(data) && data[bound] < s {
		bound *= 2
	}
	if bound > len(data)-1 {
		bound = len(data) - 1
	}
	return exponentialSearchBinaryRoutine(data, bound/2, bound, s)
}

func exponentialSearchBinaryRoutine(data []int, min, max, s int) bool {
	l := min
	r := max
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
