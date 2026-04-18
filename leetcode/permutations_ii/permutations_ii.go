package permutations_ii

import "strconv"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	res := make([][]int, 0)
	hashMap := make(map[string]bool)

	numsLeftList := permuteUnique(nums[1:])
	num := nums[0]

	for _, numsLeft := range numsLeftList {
		for i := 0; i < len(numsLeft)+1; i++ {
			arr := []int{}

			arr = append(arr, numsLeft[:i]...)
			arr = append(arr, num)
			arr = append(arr, numsLeft[i:]...)

			s := ""

			for _, n := range arr {
				s += string(strconv.Itoa(n))
			}

			if ok, _ := hashMap[s]; ok {
				continue
			}

			res = append(res, arr)
			hashMap[s] = true
		}
	}

	return res
}
