package rotateright

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	j := 1
	occurrence := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			occurrence++
		} else {
			occurrence = 1
		}

		if occurrence <= 2 {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
