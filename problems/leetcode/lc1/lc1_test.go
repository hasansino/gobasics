package lc1

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Basic example",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "Zero target",
			nums:     []int{-1, 0, 1},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "Duplicate numbers",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "Large array",
			nums:     []int{1, 5, 8, 10, 13, 21, 34, 55, 89, 144},
			target:   97,
			expected: []int{2, 8},
		},
		{
			name:     "Single solution",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "Target sum in non-adjacent elements",
			nums:     []int{4, 1, 3, 22, 5, 9, 2},
			target:   10,
			expected: []int{1, 5},
		},
		{
			name:     "All positive integers",
			nums:     []int{1, 2, 3, 4, 5},
			target:   9,
			expected: []int{3, 4},
		},
		{
			name:     "All negative integers",
			nums:     []int{-5, -4, -3, -2, -1},
			target:   -3,
			expected: []int{3, 4},
		},
		{
			name:     "No solution",
			nums:     []int{1, 2, 3},
			target:   7,
			expected: nil, // Assuming nil is returned when no solution exists
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := TwoSum(tc.nums, tc.target)

			// For cases with no expected solution
			if tc.expected == nil && result == nil {
				return
			}

			// Handle the case where a solution exists but we got nil
			if tc.expected != nil && result == nil {
				t.Errorf("TwoSum(%v, %d) = nil, expected %v", tc.nums, tc.target, tc.expected)
				return
			}

			// Handle when we expect nil but got a result
			if tc.expected == nil && result != nil {
				t.Errorf("TwoSum(%v, %d) = %v, expected nil", tc.nums, tc.target, result)
				return
			}

			// Check if result matches expected (order doesn't matter)
			if !reflect.DeepEqual(result, tc.expected) {
				// Try reverse order
				if !(len(result) == 2 && len(tc.expected) == 2 &&
					result[0] == tc.expected[1] && result[1] == tc.expected[0]) {
					t.Errorf("TwoSum(%v, %d) = %v, expected %v", tc.nums, tc.target, result, tc.expected)
				}
			}

			// Verify the solution is correct by checking the sum
			if result != nil && len(result) == 2 {
				if result[0] < 0 || result[0] >= len(tc.nums) ||
					result[1] < 0 || result[1] >= len(tc.nums) {
					t.Errorf("TwoSum(%v, %d) returned out of bounds indices: %v", tc.nums, tc.target, result)
				} else if tc.nums[result[0]]+tc.nums[result[1]] != tc.target {
					t.Errorf("TwoSum(%v, %d) = %v, but these indices don't sum to target: %d + %d = %d",
						tc.nums, tc.target, result, tc.nums[result[0]], tc.nums[result[1]], tc.nums[result[0]]+tc.nums[result[1]])
				}
			}
		})
	}
}

// TestTwoSumEmptyArray tests the behavior with an empty array
func TestTwoSumEmptyArray(t *testing.T) {
	result := TwoSum([]int{}, 0)
	if result != nil {
		t.Errorf("TwoSum([], 0) = %v, expected nil", result)
	}
}

// TestTwoSumPerformance tests with a large array to test performance
func TestTwoSumPerformance(t *testing.T) {
	// Create a large array
	largeArray := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		largeArray[i] = i
	}

	target := 19997 // Sum of 9998 + 9999
	result := TwoSum(largeArray, target)

	expected := []int{9998, 9999}
	if !reflect.DeepEqual(result, expected) && !reflect.DeepEqual(result, []int{9999, 9998}) {
		t.Errorf("TwoSum performance test failed: got %v, expected %v", result, expected)
	}
}
