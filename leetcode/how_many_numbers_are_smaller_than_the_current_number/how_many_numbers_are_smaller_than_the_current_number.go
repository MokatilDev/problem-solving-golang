package how_many_numbers_are_smaller_than_the_current_number

import "slices"

func smallerNumbersThanCurrent(nums []int) []int {
	temp := make([]int, len(nums))
	res := make([]int, len(nums))

	copy(temp, nums)
	hashMap := make(map[int]int)
	slices.Sort(temp)

	for i, n := range temp {
		if _, ok := hashMap[n]; !ok {
			hashMap[n] = i
		}
	}

	for i := range nums {
		res[i] = hashMap[nums[i]]
	}

	return res
}
