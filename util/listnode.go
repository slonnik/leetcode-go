package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeFromSlice(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	var node *ListNode
	var resultNode *ListNode

	for _, item := range data {
		if node == nil {
			node = &ListNode{Val: item}
			resultNode = node
		} else {

			node.Next = &ListNode{Val: item}
			node = node.Next
		}

	}

	return resultNode
}

func (node *ListNode) ToSlice() []int {
	result := []int{}
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func (node *ListNode) Invert() *ListNode {
	var result *ListNode
	for node != nil {
		result = &ListNode{Val: node.Val, Next: result}
		node = node.Next
	}
	return result
}
