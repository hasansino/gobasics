package sorting

import (
	"github.com/hasansino/gobasics/structures/heap"
)

// https://en.wikipedia.org/wiki/Heapsort

func HeapSort(data []int) []int {
	h := heap.NewHeap(heap.MinHeap, func(i, j interface{}) bool {
		return i.(int) <= j.(int)
	})
	for _, v := range data {
		h.Insert(v)
	}
	sorted := make([]int, 0, len(data))
	for pop := h.Pop(); pop != nil; pop = h.Pop() {
		sorted = append(sorted, pop.(int))
	}
	return sorted
}
