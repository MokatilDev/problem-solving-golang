package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			curr := root.Left
			for curr.Right != nil {
				curr = curr.Right
			}

			curr.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}
