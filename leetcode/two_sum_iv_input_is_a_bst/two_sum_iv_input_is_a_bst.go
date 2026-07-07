package two_sum_iv_input_is_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	hashMap := make(map[int]bool)
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if ok := hashMap[k-node.Val]; ok {
			return true
		}

		hashMap[node.Val] = true

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return false
}
