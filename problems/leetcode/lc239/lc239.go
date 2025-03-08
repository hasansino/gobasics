package lc239

// https://leetcode.com/problems/sliding-window-maximum/

import (
	"container/list"
)

// O(n*k) - space
// O(n+k) - memory
func maxSlidingWindowSlow(nums []int, k int) []int {
	answer := make([]int, 0, k)
	list := list.New()

	for i := 0; i < len(nums); i++ {
		list.PushFront(nums[i])
		if list.Len() > k {
			last := list.Back()
			list.Remove(last)
		}

		if list.Len() == k {
			start := true
			max := 0
			for v := list.Front(); v != nil; v = v.Next() {
				num := v.Value.(int)
				if start {
					max = v.Value.(int)
					start = false
				} else if num > max {
					max = num
				}
			}
			answer = append(answer, max)
		}
	}
	return answer
}

//| Step | `i` | Current Element (`nums[i]`) | Deque (Indices) | Deque (Values) | Window         | Max in Window |
//|------|-----|-----------------------------|-----------------|----------------|----------------|---------------|
//| 1    | 0   | 1                           | [0]             | [1]            | [1]            | -             |
//| 2    | 1   | 3                           | [1]             | [3]            | [1, 3]         | -             |
//| 3    | 2   | -1                          | [1, 2]          | [3, -1]        | [1, 3, -1]     | 3             |
//| 4    | 3   | -3                          | [1, 2, 3]       | [3, -1, -3]    | [3, -1, -3]    | 3             |
//| 5    | 4   | 5                           | [4]             | [5]            | [-1, -3, 5]    | 5             |
//| 6    | 5   | 3                           | [4, 5]          | [5, 3]         | [-3, 5, 3]     | 5             |
//| 7    | 6   | 6                           | [6]             | [6]            | [5, 3, 6]      | 6             |
//| 8    | 7   | 7                           | [7]             | [7]            | [3, 6, 7]      | 7             |

// O(n) - space
func maxSlidingWindowFast(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	answer := make([]int, 0, len(nums)-k+1)
	deque := list.New()

	//fmt.Printf("nums => %v \n", nums)

	for i := 0; i < len(nums); i++ {

		if deque.Len() > 0 && deque.Front().Value.(int) < i-k+1 {
			deque.Remove(deque.Front())
		}

		for deque.Len() > 0 && nums[deque.Back().Value.(int)] < nums[i] {
			deque.Remove(deque.Back())
		}

		deque.PushBack(i)

		if i >= k-1 {
			answer = append(answer, nums[deque.Front().Value.(int)])
		}

		//fmt.Printf("list => ")
		//for e := deque.Front(); e != nil; e = e.Next() {
		//	fmt.Printf("%v ", nums[e.Value.(int)])
		//}
		//fmt.Printf("\n")
	}

	//fmt.Printf("answer => %v \n", answer)

	return answer
}
