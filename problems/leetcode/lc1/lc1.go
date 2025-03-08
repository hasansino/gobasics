package lc1

// https://leetcode.com/problems/two-sum/description/

func TwoSum(nums []int, sum int) []int {

	mem := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := mem[nums[i]]; ok {
			return []int{mem[nums[i]], i}
		} else {
			mem[sum-nums[i]] = i
		}
	}

	return nil
}
