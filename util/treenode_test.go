package util

import (
	"reflect"
	"testing"
)

func TestTreeNode_FromSlice(t *testing.T) {

	var treeNode = TreeNodeFromSlice([]int{2, 1, 3, 4, 5, 6, 7})
	if treeNode.Val != 2 {
		t.Fatalf("Invalid root")
	}

	if treeNode.Left.Val != 1 {
		t.Fatalf("Invalid left")
	}

	if treeNode.Left.Left.Val != 4 {
		t.Fatalf("Invalid left")
	}

	if treeNode.Left.Right.Val != 5 {
		t.Fatalf("Invalid right")
	}

	if treeNode.Right.Val != 3 {
		t.Fatalf("Invalid right")

	}

	if treeNode.Right.Left.Val != 6 {
		t.Fatalf("Invalid left")

	}

	if treeNode.Right.Right.Val != 7 {
		t.Fatalf("Invalid left")

	}
}

func TestTreeNode_ToSlice(t *testing.T) {
	node := TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	slice := node.ToSlice()
	if !reflect.DeepEqual(slice, []int{2, 3, 1}) {
		t.Fatalf("Failed to make slice: %v", slice)
	}
}
