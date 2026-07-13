package max_sum_of_a_pair_with_equal_sum_of_digits

func maximumSum(nums []int) int {
	res := -1
	hashMap := make(map[int]int)

	for _, num := range nums {
		sum := sumDigites(num)

		if value, exist := hashMap[sum]; exist {
			res = max(res, value+num)
		}

		hashMap[sum] = max(hashMap[sum], num)
	}

	return res
}

func sumDigites(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}
