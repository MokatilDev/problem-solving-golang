package sorting

func selectionSort(arr *[]int) {
	nums := *arr

	for i := 0; i <= len(nums)-2; i++ {
		min_indice := i
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[min_indice] > nums[j] {
				min_indice = j
			}
		}
		nums[i], nums[min_indice] = nums[min_indice], nums[i]
	}
}

func insertionSort(arr *[]int) {
	nums := *arr
	for i := 1; i <= len(nums)-1; i++ {
		j := i - 1
		for j >= 0 && nums[j+1] < nums[j] {
			temp := nums[j+1]
			nums[j+1] = nums[j]
			nums[j] = temp
			j -= 1
		}
	}
}

func bubbleSort(arr *[]int) {
	nums := *arr
	for i := 0; i <= len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j+1] < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
