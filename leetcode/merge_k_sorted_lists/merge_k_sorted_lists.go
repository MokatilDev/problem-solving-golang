package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	first := mergeKLists(lists[:len(lists)/2])
	second := mergeKLists(lists[len(lists)/2:])

	return merge(first, second)
}

func merge(first, second *ListNode) *ListNode {
	list := &ListNode{}
	current := list

	for first != nil && second != nil {
		if first.Val < second.Val {
			current.Next = first
			first = first.Next
		} else {
			current.Next = second
			second = second.Next
		}
		current = current.Next
	}

	if first != nil {
		current.Next = first
	} else if second != nil {
		current.Next = second
	}

	return list.Next
}
