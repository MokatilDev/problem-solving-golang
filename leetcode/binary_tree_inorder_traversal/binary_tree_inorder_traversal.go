package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]int, 0)
	if root == nil {
		return nil
	}

	stack = append(stack, inorderTraversal(root.Left)...)
	stack = append(stack, root.Val)
	stack = append(stack, inorderTraversal(root.Right)...)

	return stack
}
