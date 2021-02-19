package searching

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/hasansino/gobasics/algorithms/sorting"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	data       []int
	dataSorted []int
	excluded   int // number that will be absent from data slice
}

var (
	modules = map[string]func(data, dataSorted []int, s int) bool{
		"binary": BinarySearch,
	}
)

func genRandomTestCases(num, len, max int) []testCase {
	rndCases := make([]testCase, 0, num)
	for i := 0; i < num; i++ {
		s := make([]int, 0, len)
		ex := rand.Intn(max)
		for j := 0; j < len; j++ {
			n := rand.Intn(max)
			for n == ex {
				n = rand.Intn(max)
			}
			s = append(s, n)
		}
		ss := make([]int, len)
		copy(ss, s)
		sort.Sort(sorting.SortableSlice(ss))
		rndCases = append(rndCases, testCase{data: s, dataSorted: ss, excluded: ex})
	}
	return rndCases
}

func TestSearchingModules(t *testing.T) {
	testCases := genRandomTestCases(5, 100, 999)
	for name, fn := range modules {
		t.Run(name, func(t *testing.T) {
			for _, ts := range testCases {
				assert.True(t, fn(ts.data, ts.dataSorted, ts.data[rand.Intn(len(ts.data))]))
				assert.False(t, fn(ts.data, ts.dataSorted, ts.excluded))
			}
		})
	}
}

func BenchmarkSearchingModules(b *testing.B) {
	testCases := genRandomTestCases(5, 1000, 9999)
	for name, fn := range modules {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for _, ts := range testCases {
				for n := 0; n < b.N; n++ {
					fn(ts.data, ts.dataSorted, ts.data[rand.Intn(len(ts.data))])
				}
			}
		})
	}
}

func BenchmarkSearchingModules_WorstCase(b *testing.B) {
	testCases := genRandomTestCases(5, 1000, 9999)
	for name, fn := range modules {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for _, ts := range testCases {
				for n := 0; n < b.N; n++ {
					fn(ts.data, ts.dataSorted, ts.excluded)
				}
			}
		})
	}
}
