package delete_node_in_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var prev *TreeNode
	current := root
	prev = nil

	for current != nil {
		if current.Val < key {
			prev = current
			current = current.Right
		} else if current.Val > key {
			prev = current
			current = current.Right
		} else {
			break
		}
	}

	if prev != nil {
		prev.Right = current.Right
	}

	return root
}
