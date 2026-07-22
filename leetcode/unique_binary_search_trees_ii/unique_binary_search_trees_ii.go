package unique_binary_search_trees_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return createTree(1, n)
}

func createTree(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	trees := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftTrees := createTree(start, i-1)
		rigthTrees := createTree(i+1, end)

		for left := range leftTrees {
			for right := range rigthTrees {
				root := &TreeNode{Val: i}
				root.Left = leftTrees[left]
				root.Right = rigthTrees[right]
				trees = append(trees, root)
			}
		}
	}

	return trees
}
