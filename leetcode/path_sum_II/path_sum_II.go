package path_sum_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			return append(res, []int{root.Val})
		}
		return res
	}

	for _, path := range pathSum(root.Left, targetSum-root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}

	for _, path := range pathSum(root.Right, targetSum-root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}

	return res
}
