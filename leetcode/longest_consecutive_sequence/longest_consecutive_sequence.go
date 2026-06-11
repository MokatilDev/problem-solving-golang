package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var res int

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	longest := 0
	for _, num := range nums {
		if !set[num-1] {
			curr := num
			length := 1

			for set[curr+1] {
				curr++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	res = longest

	return res
}
