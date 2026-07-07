package minimum_distance_between_bst_nodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	res, prev := math.MaxInt, -1

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if prev != -1 {
			res = min(res, node.Val-prev)
		}

		prev = node.Val

		dfs(node.Right)
	}

	dfs(root)

	return res
}
