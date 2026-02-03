package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	arr := []*TreeNode{root}

	for len(arr) > 0 {
		var values []int
		l := len(arr)

		for range l {
			node := arr[0]
			arr = arr[1:]

			values = append(values, node.Val)

			if node.Left != nil {
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}

		result = append(result, values)
	}

	return result
}
