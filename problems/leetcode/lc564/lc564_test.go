package lc564

import (
	"testing"
)

func TestNearestPalindromic(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1", "0"},
		{"15", "11"},
		{"10", "9"},
		{"21", "22"},
		{"1213", "1221"},
		{"1113", "1111"},
		{"11213", "11211"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			if result := nearestPalindromic(test.input); result != test.expected {
				t.Errorf("nearestPalindromic(%s) = %s; want %s", test.input, result, test.expected)
			}
		})
	}
}
