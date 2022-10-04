package problem226

/**
https://leetcode.com/problems/invert-binary-tree/
*/
import "github.com/slonnik/leetcode-go/util"

func InvertTree(root *util.TreeNode) *util.TreeNode {
	if root == nil {
		return root
	}

	tempNode := root.Right
	root.Right = root.Left
	root.Left = tempNode
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}
