package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CustomNode struct {
	Node *TreeNode
	Min  *int
	Max  *int
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []CustomNode{{Node: root, Min: nil, Max: nil}}

	for len(queue) > 0 {
		l := len(queue)

		for range l {
			current := queue[0]
			queue = queue[1:]

			val := current.Node.Val

			if (current.Min != nil && val <= *current.Min) ||
				(current.Max != nil && val >= *current.Max) {
				return false
			}

			if current.Node.Left != nil {
				queue = append(queue, CustomNode{Node: current.Node.Left, Min: current.Min, Max: &val})
			}

			if current.Node.Right != nil {
				queue = append(queue, CustomNode{Node: current.Node.Right, Min: &val, Max: current.Max})
			}
		}
	}

	return true
}
