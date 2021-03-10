package searchnum

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	data   string
	substr string
}

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	modules = map[string]func(data string, s string) int{
		"naive":     Naive,
		"rabinkarp": RabinKarp,
	}
)

func genRandomTestCases(num, l int) []testCase {
	rndCases := make([]testCase, 0, num)
	for i := 0; i < num; i++ {
		var s []rune
		for j := 0; j < l; j++ {
			s = append(s, letters[rand.Intn(len(letters))])
		}
		rndCases = append(rndCases, testCase{
			data:   string(s),
			substr: string(s[len(s)/2 : len(s)/2+rand.Intn(25)+1])},
		)
	}
	return rndCases
}

func TestSearchingModules(t *testing.T) {
	for name, fn := range modules {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, -1, fn("", ""))
		})
	}
	testCases := genRandomTestCases(5, 1000)
	for name, fn := range modules {
		t.Run(name, func(t *testing.T) {
			for _, ts := range testCases {
				assert.Greater(t, fn(ts.data, ts.substr), -1)
				assert.Equal(t, -1, fn(ts.data, "!@#$%^&"))
			}
		})
	}
}

func BenchmarkSearchingModules(b *testing.B) {
	testCases := genRandomTestCases(5, 1000)
	for name, fn := range modules {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for _, ts := range testCases {
				for n := 0; n < b.N; n++ {
					fn(ts.data, ts.substr)
				}
			}
		})
	}
}

func BenchmarkSearchingModules_WorstCase(b *testing.B) {
	testCases := genRandomTestCases(5, 1000)
	for name, fn := range modules {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for _, ts := range testCases {
				for n := 0; n < b.N; n++ {
					fn(ts.data, "#$%*)(@&^")
				}
			}
		})
	}
}
