package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeNodeFromSlice(data []int) *TreeNode {

	node := &TreeNode{Val: data[0]}
	internalNodeFromSlice(node, 1, data)
	return node
}

func internalNodeFromSlice(rootNode *TreeNode, rootIndex int, data []int) {

	leftIndex := 2 * rootIndex
	rightIndex := leftIndex + 1

	if rightIndex > len(data) {
		return
	}

	rootNode.Left = &TreeNode{Val: data[leftIndex-1]}
	rootNode.Right = &TreeNode{Val: data[rightIndex-1]}

	internalNodeFromSlice(rootNode.Left, leftIndex, data)
	internalNodeFromSlice(rootNode.Right, rightIndex, data)
}

func (node *TreeNode) ToSlice() []int {
	data := []int{}
	node.internalToSlice(data)
	return data
}

func (node *TreeNode) internalToSlice(data []int) {
	if node == nil {
		return
	}
	node.Left.internalToSlice(data)
	node.Right.internalToSlice(data)
}
