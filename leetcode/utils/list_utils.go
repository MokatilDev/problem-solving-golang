package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinkList(list []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, val := range list {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

func IsEqaul(list1, list2 *ListNode) bool {
	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	return list1 == nil && list2 == nil
}

func PrintList(list *ListNode) []int {
	var res []int
	current := list
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	return res
}
