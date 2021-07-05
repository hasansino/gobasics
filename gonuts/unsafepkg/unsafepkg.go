//
// Package unsafepkg demonstrates usage of unsafe package.
//
package unsafepkg

import "unsafe"

// B2i converts bool value to int8 (or uint8)
// Since both types share same size (1 byte), it is possible
// to convert them using unsafe approach.
//
// Difference in performance is almost unnoticeable in such simple case.
//
// goos: linux
// goarch: amd64
// pkg: github.com/hasansino/gobasics/gonuts/unsafepkg
// BenchmarkB2i-8                  1000000000               0.306 ns/op
// BenchmarkB2iStandard-8          1000000000               0.307 ns/op
//
func B2i(b bool) int8 {
	// (unsafe.Pointer(&b) - create new unsafe pointer with address of b
	// (*int8) - convert to pointer of type int8
	// * - take value of last pointer (int8)
	return *(*int8)(unsafe.Pointer(&b))
}

// _B2i is traditional way of conversion using new variable
func _B2i(b bool) int8 {
	var i int8
	if b {
		i = 1
	}
	return i
}
