//
// Package mapkeyasstruct is performance comparison between arbitrary structs
// as a indexes for maps and ordinary strings. Any value used as index needs
// to be transformed to hash before use and larger / complex is the data set
// - the longer time it takes.
//
//  goos: linux
//  goarch: amd64
//  pkg: github.com/hasansino/gobasics/gonuts/mapkeyasstruct
//  cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
//  BenchmarkStringKey-8   	39877777	        29.94 ns/op
//  BenchmarkStructKey-8   	23088190	        48.09 ns/op

package mapkeyasstruct

type A map[string]int

type key struct {
	fieldA string
	fieldB int32
}

type B map[key]int
