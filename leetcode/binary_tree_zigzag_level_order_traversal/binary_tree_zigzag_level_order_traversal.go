package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		arr := make([]int, 0, l)

		for range l {

			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if len(res)%2 == 0 {
				arr = append(arr, node.Val)
			} else {
				arr = append([]int{node.Val}, arr...)
			}
		}

		res = append(res, arr)
	}

	return res
}
