package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	pivot := 0
	for pivot < len(postorder) && inorder[pivot] != postorder[len(postorder)-1] {
		pivot++
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	root.Val = postorder[len(postorder)-1]
	root.Left = buildTree(inorder[:pivot], postorder[:pivot])
	root.Right = buildTree(inorder[pivot+1:], postorder[pivot:len(postorder)-1])

	return root
}
