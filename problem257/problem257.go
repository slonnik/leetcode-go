package problem257

import (
	"fmt"
	"github.com/slonnik/leetcode-go/util"
)

//https://leetcode.com/problems/binary-tree-paths/

var data []string

func binaryTreePaths(root *util.TreeNode) []string {
	path := ""
	data = []string{}
	internal(root, path)
	return data
}

func internal(node *util.TreeNode, path string) {
	if node == nil {
		return
	}
	if path == "" {
		path = fmt.Sprintf("%v", node.Val)
	} else {
		path = fmt.Sprintf("%v->%v", path, node.Val)
	}

	if node.Left == nil && node.Right == nil {
		data = append(data, path)
		return
	}
	internal(node.Left, path)
	internal(node.Right, path)
}
