package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)

		for i := range l {
			current := queue[0]
			queue = queue[1:]

			if i == 0 {
				res = append(res, current.Val)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
		}
	}

	return res
}
