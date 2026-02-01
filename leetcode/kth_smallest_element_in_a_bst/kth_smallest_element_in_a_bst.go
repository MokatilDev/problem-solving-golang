package kth_smallest_elementina_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode

	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return node.Val
		}

		current = node.Right
	}

	return -1
}
