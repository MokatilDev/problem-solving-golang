package find_the_duplicate_number

func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	fast := nums[0]
	slow := nums[0]

	for {
		fast = nums[nums[fast]]
		slow = nums[slow]

		if fast == slow {
			break
		}
	}

	slow = nums[0]

	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]

		if fast == slow {
			break
		}
	}

	return slow
}
