package sum_of_root_to_leaf_binary_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return pathNumber(root, 0)
}

func pathNumber(node *TreeNode, current int) int {
	if node == nil {
		return 0
	}

	current = current*2 + node.Val

	if node.Left == nil && node.Right == nil {
		return current
	}

	left_sum := pathNumber(node.Left, current)
	right_sum := pathNumber(node.Right, current)

	return left_sum + right_sum
}
