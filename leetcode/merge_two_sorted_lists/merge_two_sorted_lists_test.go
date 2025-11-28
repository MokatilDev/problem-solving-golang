package mergeTwoSortedLists

import (
	"testing"

	"github.com/MokatilDev/Problem-Solving-Lab/leetcode/utils"
)

func TestMergeTwoLists(t *testing.T) {
	type testCase struct {
		name     string
		input1   []int
		input2   []int
		expected []int
	}

	tests := []testCase{
		{
			name:     "Standard Case",
			input1:   []int{1, 2, 4},
			input2:   []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Empty Lists",
			input1:   []int{},
			input2:   []int{},
			expected: []int{},
		},
		{
			name:     "One Empty List",
			input1:   []int{},
			input2:   []int{0},
			expected: []int{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list1 := utils.BuildLinkList(tc.input1)
			list2 := utils.BuildLinkList(tc.input2)

			expected := utils.BuildLinkList(tc.expected)
			res := mergeTwoLists(list1, list2)

			if !utils.IsEqaul(expected, res) {
				t.Error("Not Passed check the code")
			}
		})
	}
}
