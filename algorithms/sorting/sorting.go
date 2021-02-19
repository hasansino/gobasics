package sorting

type SortableSlice []int

func (s SortableSlice) Len() int           { return len(s) }
func (s SortableSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s SortableSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
