package average_of_levels_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	res := []float64{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}
		n := len(queue)
		for range n {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		sum := 0
		for _, v := range level {
			sum += v
		}

		res = append(res, float64(sum)/float64(len(level)))
	}

	return res
}
