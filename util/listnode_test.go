package util

import (
	"reflect"
	"testing"
)

func TestListNode_FromSlice(t *testing.T) {
	node := ListNodeFromSlice([]int{1, 2, 3})
	if node.Val != 1 || node.Next == nil {
		t.Fatalf("Invalid node %v", node)
	}
	node = node.Next
	if node.Val != 2 || node.Next == nil {
		t.Fatalf("Invalid node %v", node)
	}

	node = node.Next
	if node.Val != 3 || node.Next != nil {
		t.Fatalf("Invalid node %v", node)
	}
}

func TestListNode_ToSlice(t *testing.T) {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	data := node.ToSlice()

	if !reflect.DeepEqual(data, []int{1, 2, 3}) {
		t.Fatalf("Invalid data %v", data)
	}
}

func TestListNode_Invert(t *testing.T) {
	node := ListNodeFromSlice([]int{1, 2, 3})
	data := node.Invert().ToSlice()

	if !reflect.DeepEqual(data, []int{3, 2, 1}) {
		t.Fatalf("Invalid data %v", data)
	}
}
