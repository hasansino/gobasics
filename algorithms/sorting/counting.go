package sorting

// https://en.wikipedia.org/wiki/Counting_sort

// CountingSort with negative integer support
func CountingSort(data []int) []int {
	min, max := countingSortMinMaxNum(data)
	if max == 0 {
		max = 1
	}

	var offset int
	if min < 0 {
		offset = min * -1
		max += offset
	}

	counters := make([]int, max+1)
	for i := 0; i < len(data); i++ {
		counters[data[i]+offset]++
	}

	tmp := make([]int, 0)
	for k, v := range counters {
		if v == 0 {
			continue
		}
		for j := 0; j < v; j++ {
			tmp = append(tmp, k-offset)
		}
	}

	return tmp
}

func countingSortMinMaxNum(data []int) (min, max int) {
	for i := 0; i < len(data); i++ {
		if i == 0 || data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}
	return
}
