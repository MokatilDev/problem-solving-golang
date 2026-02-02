package construct_binary_tree_from_preorder_and_inorder_traversal

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i := slices.Index(inorder, preorder[0])
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}
