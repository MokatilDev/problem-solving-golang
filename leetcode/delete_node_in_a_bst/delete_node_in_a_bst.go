package delete_node_in_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := minValueNode(root.Right)
			root.Val = minNode.Val
			root.Right = deleteNode(root.Right, minNode.Val)
		}
	}

	return root
}

func minValueNode(root *TreeNode) *TreeNode {
	current := root

	for current != nil && current.Left != nil {
		current = current.Left
	}

	return current
}
