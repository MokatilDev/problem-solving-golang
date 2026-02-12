package sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return pathNumber(root, 0)
}

func pathNumber(node *TreeNode, current int) int {
	if node == nil {
		return 0
	}

	newCurrent := current*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return newCurrent
	}

	left_sum := pathNumber(node.Left, 0)
	right_sum := pathNumber(node.Right, 0)

	return left_sum + right_sum
}
