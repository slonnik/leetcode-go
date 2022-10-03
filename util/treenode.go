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
	count := node.countNodes()
	data := make([]int, count)
	node.internalToSlice(1, data)
	return data
}

func (node *TreeNode) countNodes() int {
	if node == nil {
		return 0
	}
	count := 1
	count += node.Left.countNodes()
	count += node.Right.countNodes()
	return count
}

func (node *TreeNode) internalToSlice(rootIndex int, data []int) {
	if node == nil {
		return
	}
	data[rootIndex-1] = node.Val
	node.Left.internalToSlice(2*rootIndex, data)
	node.Right.internalToSlice(2*rootIndex+1, data)
}
