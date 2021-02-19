package sorting

// https://en.wikipedia.org/wiki/Radix_sort

func RadixSort(data []int) []int {
	maxNumLen := radixSortMaxNumLen(data) // how many digits is in largest number?
	for i := 1; i <= maxNumLen; i++ {     // for each digit position
		data = radixSortCountingSubSort(data, i)
	}
	return data
}

func radixSortCountingSubSort(data []int, d int) []int {
	var (
		sorted    = make([]int, 0, len(data))
		counters  = make([][]int, 10) // base10
		nCounters = make([][]int, 10) // for negative numbers
	)
	for j := 0; j < len(data); j++ {
		n := radixSortSignificantDigitValue(data[j], d)
		if n < 0 {
			nCounters[n*-1] = append(nCounters[n*-1], data[j])
		} else {
			counters[n] = append(counters[n], data[j])
		}
	}
	// negative numbers first
	for i := len(nCounters) - 1; i > 0; i-- {
		sorted = append(sorted, nCounters[i]...)
	}
	for i := 0; i < len(counters); i++ {
		sorted = append(sorted, counters[i]...)
	}
	return sorted
}

func radixSortMaxNumLen(data []int) int {
	var max int
	for _, v := range data {
		if v < 0 {
			v *= -1 // take absolute value
		}
		if v > max {
			max = v
		}
	}
	var l int
	for max != 0 {
		max /= 10
		l++
	}
	return l
}

func radixSortSignificantDigitValue(n, pos int) int {
	var digit int
	for pos > 0 {
		digit = n % 10
		n /= 10
		pos--
	}
	return digit
}
