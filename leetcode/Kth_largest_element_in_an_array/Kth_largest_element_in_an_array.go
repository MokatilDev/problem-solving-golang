package Kth_largest_element_in_an_array

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	values := *h
	n := len(values)

	x := values[n-1]
	*h = values[:n-1]

	return x
}

func findKthLargest(nums []int, k int) int {
	values := make(MinHeap, k)
	copy(values, nums[:k])

	heap.Init(&values)

	for _, num := range nums[k:] {
		if num > values[0] {
			heap.Pop(&values)
			heap.Push(&values, num)
		}
	}

	return values[0]
}
