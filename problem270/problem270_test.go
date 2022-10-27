package problem270

import (
	"github.com/slonnik/leetcode-go/util"
	"testing"
)

func TestCaseOne(t *testing.T) {
	node := &util.TreeNode{Val: 4, Right: &util.TreeNode{Val: 5}, Left: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 1}, Right: &util.TreeNode{Val: 3}}}
	result := closestValue(node, 3.714286)
	if result != 4 {
		t.Errorf("Expecting 4, but we have %v", result)
	}
}

func TestCaseTwo(t *testing.T) {
	node := &util.TreeNode{Val: 1}
	result := closestValue(node, 4.421)
	if result != 1 {
		t.Errorf("Expecting 1, but we have %v", result)
	}
}

func TestCaseThree(t *testing.T) {
	node := &util.TreeNode{Val: 1, Right: &util.TreeNode{Val: 2}}
	result := closestValue(node, 3.428571)
	if result != 2 {
		t.Errorf("Expecting 1, but we have %v", result)
	}
}
