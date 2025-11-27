package mergetwosortedlists

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type testCase struct {
		input1   []int
		input2   []int
		expected []int
	}

	tests := []testCase{
		{
			input1:   []int{1, 2, 4},
			input2:   []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input1:   []int{1, 2, 4},
			input2:   []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input1:   []int{1, 2, 4},
			input2:   []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
	}

	for _, tc := range tests {
		t.Run("linked list", func(t *testing.T) {
			input1 := tc.input1
			input2 := tc.input2
			list1 := buildLinkList(input1)
			list2 := buildLinkList(input2)
			expected := buildLinkList(tc.expected)

			if !isEqaul(expected, mergeTwoLists(list1, list2)) {
				t.Error("Not Passed check the code")
			}
		})
	}
}

func buildLinkList(list []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, val := range list {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

func isEqaul(list1, list2 *ListNode) bool {
	for list1.Next != nil && list2.Next != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	return list1 == nil && list2 == nil
}

func printList(list *ListNode) []int {
	var res []int
	current := list
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	return res
}
