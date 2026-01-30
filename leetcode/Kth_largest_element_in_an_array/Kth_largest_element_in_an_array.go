package Kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	return mergeSort(nums)[len(nums)-k]
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	first := mergeSort(nums[:len(nums)/2])
	second := mergeSort(nums[len(nums)/2:])

	return merge(first, second)
}

func merge(a, b []int) []int {
	arr := make([]int, 0, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			arr = append(arr, a[i])
			i++
		} else {
			arr = append(arr, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		arr = append(arr, a[i])
	}

	for ; j < len(b); j++ {
		arr = append(arr, b[j])
	}

	return arr
}
