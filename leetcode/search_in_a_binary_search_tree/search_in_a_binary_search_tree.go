package search_in_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	current := root

	for current != nil {
		if current.Val < val {
			current = current.Right
		} else if current.Val > val {
			current = current.Left
		} else {
			return current
		}
	}

	return nil
}
