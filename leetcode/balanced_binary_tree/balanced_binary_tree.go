package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		if left == -1 {
			return -1
		}

		right := dfs(node.Right)
		if right == -1 {
			return -1
		}

		diff := left - right
		if diff > 1 || diff < -1 {
			return -1
		}

		return max(left, right) + 1
	}

	return dfs(root) != -1
}
