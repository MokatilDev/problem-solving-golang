package linked_list_cycle_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}
