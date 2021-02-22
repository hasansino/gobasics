package searching

import "math"

// https://en.wikipedia.org/wiki/Jump_search

func JumpSearch(data []int, s int) bool {
	if len(data) == 0 {
		return false
	}
	var (
		pos       int
		maxIdx    = len(data) - 1
		blockSize = int(math.Sqrt(float64(len(data))))
	)
	for data[pos] < s {
		pos += blockSize
		if pos > maxIdx {
			pos = maxIdx
		}
		if pos == maxIdx {
			break
		}
	}
	if data[pos] < s {
		return false
	}
	// backtrack
	for data[pos] > s {
		if pos == 0 {
			break
		}
		pos--
	}
	return data[pos] == s
}
