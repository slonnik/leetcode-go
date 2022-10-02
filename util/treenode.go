package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeNodeFromSlice(data []int) *TreeNode {

	node := &TreeNode{Val: data[0]}
	internalNode(node, 1, data)
	return node
}

func internalNode(rootNode *TreeNode, rootIndex int, data []int) {

	leftIndex := 2 * rootIndex
	rightIndex := leftIndex + 1

	if rightIndex > len(data) {
		return
	}

	rootNode.Left = &TreeNode{Val: data[leftIndex-1]}
	rootNode.Right = &TreeNode{Val: data[rightIndex-1]}

	internalNode(rootNode.Left, leftIndex, data)
	internalNode(rootNode.Right, rightIndex, data)
}
