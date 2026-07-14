package linked_list_components

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}

	count := 0
	current := head
	inComponent := false

	for current != nil {
		if hashMap[current.Val] {
			if !inComponent {
				count++
				inComponent = true
			}
		} else {
			inComponent = false
		}

		current = current.Next
	}

	return count
}
