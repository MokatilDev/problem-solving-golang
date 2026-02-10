package maximum_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	rootVal := 0
	rootIndex := 0

	for i, num := range nums {
		if rootVal < num {
			rootVal = num
			rootIndex = i
		}
	}

	root := &TreeNode{}

	root.Val = rootVal
	root.Left = constructMaximumBinaryTree(nums[:rootIndex])
	root.Right = constructMaximumBinaryTree(nums[rootIndex+1:])

	return root
}
