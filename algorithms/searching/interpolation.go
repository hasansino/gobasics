package searching

// https://en.wikipedia.org/wiki/Interpolation_search

func InterpolationSearch(data []int, s int) bool {
	if len(data) == 0 {
		return false
	}
	var (
		min, mid int
		max      = len(data) - 1
	)
	for data[max] != data[min] && s >= data[min] && s <= data[max] {
		mid = min + ((s - data[min]) * (max - min) / (data[max] - data[min]))

		switch {
		case s > data[mid]:
			min = mid + 1
		case s < data[mid]:
			max = mid - 1
		default:
			return true
		}
	}
	return s == data[min]
}
