package find_largest_value_in_each_tree_row

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currMax := math.MinInt

		l := len(queue)

		for range l {
			node := queue[0]
			queue = queue[1:]

			if currMax < node.Val {
				currMax = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, currMax)
	}

	return res
}
