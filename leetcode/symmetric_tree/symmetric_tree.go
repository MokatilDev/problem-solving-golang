package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		leftNode := queue[0]
		rightNode := queue[1]

		queue = queue[2:]

		if leftNode == nil && rightNode == nil {
			continue
		}

		if leftNode == nil || rightNode == nil {
			return false
		}

		if leftNode.Val != rightNode.Val {
			return false
		}

		queue = append(queue, leftNode.Left, rightNode.Right, leftNode.Right, rightNode.Left)
	}

	return true
}
