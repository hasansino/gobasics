package searchnum

// https://en.wikipedia.org/wiki/Ternary_search

func TernarySearch(data []int, s int) bool {
	if len(data) == 0 {
		return false
	}

	var (
		left       int
		right      = len(data) - 1
		mid1, mid2 int
	)

	for left <= right {
		mid1 = left + (right-left)/3
		mid2 = right - (right-left)/3

		switch {
		// lucky shot
		case data[mid1] == s || data[mid2] == s: // lucky shot
			return true
		case data[mid1] > s: // check indexes (0->mid1)
			right = mid1 - 1
		case data[mid2] < s: // check indexes (mid2->len(data))
			left = mid2 + 1
		default: // check middle section (mid1->mid2)
			left, right = mid1+1, mid2-1
		}

		if mid1 >= mid2 {
			break
		}
	}

	return false
}
