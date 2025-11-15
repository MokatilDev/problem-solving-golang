package sorting

func selection(arr *[]int) {
	nums := *arr

	for i := 0; i <= len(nums)-2; i++ {
		temp := nums[i]
		min_indice := i
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[min_indice] > nums[j] {
				min_indice = j
			}
		}
		nums[i] = nums[min_indice]
		nums[min_indice] = temp
	}
}
