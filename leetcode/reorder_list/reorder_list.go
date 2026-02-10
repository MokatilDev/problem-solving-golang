package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	stack := []*ListNode{}
	current := head

	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	current = head
	n := len(stack)

	for i := 0; i < n/2; i++ {
		node := stack[n-1-i]

		temp := current.Next
		current.Next = node
		node.Next = temp

		current = temp
	}

	if current != nil {
		current.Next = nil
	}
}
