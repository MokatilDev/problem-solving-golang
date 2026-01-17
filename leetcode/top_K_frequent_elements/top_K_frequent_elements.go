package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)

	frequents := make(map[int]int)
	temp := make([][]int, len(nums)+1)

	for _, num := range nums {
		frequents[num]++
	}

	for key, value := range frequents {
		temp[value] = append(temp[value], key)
	}

	for i := len(temp) - 1; i >= 0 && len(res) < k; i-- {
		if len(temp[i]) > 0 {
			res = append(res, temp[i]...)
		}
	}

	return res[:k]
}
