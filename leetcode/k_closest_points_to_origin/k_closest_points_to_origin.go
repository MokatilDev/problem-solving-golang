package k_closest_points_to_origin

import "container/heap"

type MaxHeap [][]int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return distance(h[i]) > distance(h[j])
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() any {
	values := *h
	n := len(values)

	x := values[n-1]
	*h = values[:n-1]

	return x
}

func distance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func kClosest(points [][]int, k int) [][]int {
	res := MaxHeap{}
	for _, point := range points {
		heap.Push(&res, point)
		if len(res) > k {
			heap.Pop(&res)
		}
	}

	return res
}
