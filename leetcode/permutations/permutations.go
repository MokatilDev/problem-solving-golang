package permutations

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	res := make([][]int, 0)

	numsLeftList := permute(nums[1:])
	num := nums[0]

	for _, numsLeft := range numsLeftList {
		for j := 0; j < len(numsLeft)+1; j++ {
			arr := []int{}

			arr = append(arr, numsLeft[:j]...)
			arr = append(arr, num)
			arr = append(arr, numsLeft[j:]...)

			res = append(res, arr)
		}
	}

	return res
}
