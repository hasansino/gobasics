package searchnum

// https://en.wikipedia.org/wiki/String-searching_algorithm#Na%C3%AFve_string_search

func Naive(data string, s string) int {
	if s == "" {
		return -1
	}
	for j := 0; j+len(s) <= len(data); j++ {
		if data[j:j+len(s)] == s {
			return j
		}
	}
	return -1
}
