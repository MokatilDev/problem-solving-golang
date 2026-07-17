package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	left := &ListNode{}
	right := &ListNode{}

	currLeft := left
	currRight := right

	curr := head

	for curr != nil {
		if curr.Val < x {
			currLeft.Next = curr
			currLeft = currLeft.Next
		} else {
			currRight.Next = curr
			currRight = currRight.Next
		}

		curr = curr.Next
	}

	currRight.Next = nil
	currLeft.Next = right.Next

	return left.Next
}
