package binary_tree_paths

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}

	if root == nil {
		return res
	}

	var dfs func(currNode *TreeNode, path string)

	dfs = func(currNode *TreeNode, path string) {
		if currNode == nil {
			return
		}

		str := strconv.Itoa(currNode.Val)
		path += str

		if currNode.Left == nil && currNode.Right == nil {
			res = append(res, path)
		}

		path += "->"

		dfs(currNode.Left, path)
		dfs(currNode.Right, path)
	}

	dfs(root, "")

	return res
}
