package recoverbinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var node1, node2, prev *TreeNode

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		if prev != nil && prev.Val > node.Val {
			if node1 == nil {
				node1 = prev
			}

			node2 = node
		}

		prev = node
		inorder(node.Right)
	}

	inorder(root)

	if node1 != nil && node2 != nil {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
}
