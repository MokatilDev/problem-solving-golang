package balance_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var res *TreeNode
	values := make([]int, 0)

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		values = append(values, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	var buildBST func(l, r int) *TreeNode
	buildBST = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		mid := l + (r-l)/2
		root := &TreeNode{Val: values[mid]}

		root.Left = buildBST(l, mid-1)
		root.Right = buildBST(mid+1, r)

		return root
	}

	res = buildBST(0, len(values)-1)
	return res
}
