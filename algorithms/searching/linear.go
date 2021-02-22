package searching

// https://en.wikipedia.org/wiki/Linear_search

func LinearSearch(data []int, s int) bool {
	for _, v := range data {
		if v == s {
			return true
		}
	}
	return false
}
