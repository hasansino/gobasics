package sorting

type sortableSlice []int

func (s sortableSlice) Len() int           { return len(s) }
func (s sortableSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s sortableSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
