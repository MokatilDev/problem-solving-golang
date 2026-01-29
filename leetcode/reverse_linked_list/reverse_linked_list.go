package reverse_linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// Using Loop Methode
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	if head == nil {
		return nil
	}

	prev = nil
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// Useing Recursive Methode

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	revHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return revHead
}
