package util

import (
	"testing"
)

func TestCreateTreeNodeFromSlice(t *testing.T) {

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
