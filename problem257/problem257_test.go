package problem257

import (
	"github.com/slonnik/leetcode-go/util"
	"reflect"
	"sort"
	"testing"
)

func TestCaseOne(t *testing.T) {
	node := util.TreeNode{Val: 1,
		Left:  &util.TreeNode{Val: 2, Right: &util.TreeNode{Val: 5}},
		Right: &util.TreeNode{Val: 3}}

	result := binaryTreePaths(&node)
	expected := []string{"1->2->5", "1->3"}
	sort.Strings(result)
	sort.Strings(expected)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, but have %v", expected, result)
	}
}

func TestCaseTwo(t *testing.T) {
	node := util.TreeNode{Val: 1}

	result := binaryTreePaths(&node)
	expected := []string{"1"}
	sort.Strings(result)
	sort.Strings(expected)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, but have %v", expected, result)
	}
}
