package split_linked_list_in_parts

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var n int
	current := head

	for current != nil {
		n++
		current = current.Next
	}

	var res []*ListNode = make([]*ListNode, k)

	len0, r := n/k, n%k
	current = head

	for i := range k {
		res[i] = current
		extra := 0
		if r > 0 {
			extra = 1
		}

		for j := 0; j < len0-1+extra; j++ {
			if current == nil {
				break
			}
			current = current.Next
		}

		if current != nil {
			temp := current.Next
			current.Next = nil
			current = temp
		}
		r--
	}

	return res
}
