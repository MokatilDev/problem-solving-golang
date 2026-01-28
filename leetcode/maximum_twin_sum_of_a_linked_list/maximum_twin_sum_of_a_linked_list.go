package maximum_twin_sum_of_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution With O(1) Space Complexity (Two Pointer Method)

func pairSum(head *ListNode) int {
	var res int
	fast, slow := head, head
	var prev *ListNode = nil

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	for slow != nil {
		sum := slow.Val + prev.Val
		if res < sum {
			res = sum
		}
		slow = slow.Next
		prev = prev.Next
	}

	return res
}

// Solutuion With O(n) Space Complexety (Stack)

/*
func pairSum(head *ListNode) int {
	var res int
	var n int
	var stack []int

	current := head

	for current != nil {
		n++
		stack = append(stack, current.Val)
		current = current.Next
	}

	current = head

	for i := 0; i <= (n/2)-1; i++ {
		sum := stack[i] + stack[n-1-i]
		if res < sum {
			res = sum
		}
	}

	return res
}
*/
