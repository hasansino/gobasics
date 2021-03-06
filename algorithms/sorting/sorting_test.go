package sorting

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	testSlice       SortableSlice
	testSliceSorted SortableSlice
}

var (
	defTestCases = []testCase{
		{testSlice: SortableSlice{}, testSliceSorted: SortableSlice{}},
		{testSlice: SortableSlice{0}, testSliceSorted: SortableSlice{0}},
		{testSlice: SortableSlice{1, 0}, testSliceSorted: SortableSlice{0, 1}},
		{testSlice: SortableSlice{2, -2, 3, -4, 0}, testSliceSorted: SortableSlice{-4, -2, 0, 2, 3}},
		{testSlice: SortableSlice{0, 0, 0, 0, 0, 0}, testSliceSorted: SortableSlice{0, 0, 0, 0, 0, 0}},
		{
			testSlice:       SortableSlice{645, 240, 341, 899, 119, 251, 149, 587, 29, 66, 738, 858, 756, 699, 893, 746, 762, 233, 249, 194, 177, 944, 175, 81, 512, 362, 92, 811, 445, 966, 747, 753, 154, 610, 718, 61, 257, 552, 863, 428, 989, 143, 961, 366, 969, 686, 223, 663, 944, 750, 502, 225, 278, 908, 644, 696, 289, 478, 355, 149, 236, 359, 628, 728, 837, 120, 973, 637, 136, 257, 42, 537, 780, 234, 282, 16, 938, 769, 869, 67, 668, 119, 982, 970, 787, 896, 409, 547, 832, 150, 140, 489, 936, 234, 965, 33, 82, 559, 148, 18},
			testSliceSorted: SortableSlice{16, 18, 29, 33, 42, 61, 66, 67, 81, 82, 92, 119, 119, 120, 136, 140, 143, 148, 149, 149, 150, 154, 175, 177, 194, 223, 225, 233, 234, 234, 236, 240, 249, 251, 257, 257, 278, 282, 289, 341, 355, 359, 362, 366, 409, 428, 445, 478, 489, 502, 512, 537, 547, 552, 559, 587, 610, 628, 637, 644, 645, 663, 668, 686, 696, 699, 718, 728, 738, 746, 747, 750, 753, 756, 762, 769, 780, 787, 811, 832, 837, 858, 863, 869, 893, 896, 899, 908, 936, 938, 944, 944, 961, 965, 966, 969, 970, 973, 982, 989},
		},
	}
	// using standard library interface sort.Interface
	modules = map[string]func(sort.Interface){
		"sort.Sort":   sort.Sort,
		"sort.Stable": sort.Stable,
		"quick":       QuickSort,
		"bubble":      BubbleSort,
		"cocktail":    CocktailSort,
		"comb":        CombSort,
		"gnome":       GnomeSort,
		"insertion":   InsertionSort,
		"selection":   SelectionSort,
		"odd-even":    OddEvenSort,
		"shell":       ShellSort,
	}
	// using standard slice of integers
	modulesRaw = map[string]func([]int) []int{
		"merge":    MergeSort,
		"tree":     TreeSort,
		"counting": CountingSort,
		"heap":     HeapSort,
		"radix":    RadixSort,
	}
)

func genRandomTestCases(num, len, max int) []testCase {
	rndCases := make([]testCase, num)
	for i := 0; i < num; i++ {
		s := make(SortableSlice, len)
		for j := 0; j < len; j++ {
			s = append(s, rand.Intn(max))
		}
		ss := cloneUnsortedSlice(s)
		sort.Sort(ss)
		rndCases = append(rndCases, testCase{testSlice: s, testSliceSorted: ss})
	}
	return rndCases
}

func cloneUnsortedSlice(s SortableSlice) SortableSlice {
	clone := make(SortableSlice, s.Len())
	copy(clone, s)
	return clone
}

func TestSortingModules(t *testing.T) {
	for name, fn := range modules {
		t.Run(name, func(t *testing.T) {
			for _, ts := range defTestCases {
				s := cloneUnsortedSlice(ts.testSlice)
				fn(s)
				assert.Equal(t, ts.testSliceSorted, s)
			}
		})
	}
	for name, fn := range modulesRaw {
		t.Run(name, func(t *testing.T) {
			for _, ts := range defTestCases {
				s := cloneUnsortedSlice(ts.testSlice)
				assert.Equal(t, []int(ts.testSliceSorted), fn(s))
			}
		})
	}
}

func BenchmarkSortingModules_Small(b *testing.B) {
	benchmarkSortingModules(genRandomTestCases(5, 25, 99), b)
}

func BenchmarkSortingModules_Medium(b *testing.B) {
	benchmarkSortingModules(genRandomTestCases(5, 100, 999), b)
}

func BenchmarkSortingModules_Big(b *testing.B) {
	benchmarkSortingModules(genRandomTestCases(5, 1000, 9999), b)
}

func benchmarkSortingModules(benchCases []testCase, b *testing.B) {
	for name, fn := range modules {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for _, ts := range benchCases {
				for n := 0; n < b.N; n++ {
					s := cloneUnsortedSlice(ts.testSlice)
					fn(s)
				}
			}
		})
	}
	for name, fn := range modulesRaw {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for _, ts := range benchCases {
				for n := 0; n < b.N; n++ {
					s := cloneUnsortedSlice(ts.testSlice)
					fn(s)
				}
			}
		})
	}
}
