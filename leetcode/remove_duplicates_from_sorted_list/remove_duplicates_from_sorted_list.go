package removeDuplicatesFromSortedList

import "github.com/MokatilDev/problem-solving-golang/leetcode/utils"

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
