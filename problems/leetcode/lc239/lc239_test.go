package lc239

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "1",
			nums:     []int{1, 3, 1, 2, 0, 5},
			k:        3,
			expected: []int{3, 3, 2, 5},
		},
		{
			name:     "Basic example",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{8, 7, 6, 5, 4, 3, 2, 1},
			k:        3,
			expected: []int{8, 7, 6, 5, 4, 3},
		},
		{
			name:     "Single element window",
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Window size equals array size",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{5},
		},
		{
			name:     "All elements are the same",
			nums:     []int{2, 2, 2, 2, 2},
			k:        3,
			expected: []int{2, 2, 2},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "Window size larger than array",
			nums:     []int{1, 2},
			k:        3,
			expected: []int{},
		},
		{
			name:     "2",
			nums:     []int{1, 2},
			k:        2,
			expected: []int{2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSlidingWindowFast(tc.nums, tc.k)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, expected %v", tc.nums, tc.k, result, tc.expected)
			}
		})
	}
}
