package problem226

import (
	"github.com/slonnik/leetcode-go/util"
	"reflect"
	"testing"
)

func TestInvertTrivial(t *testing.T) {
	node := util.TreeNodeFromSlice([]int{2, 1, 3})
	result := InvertTree(node).ToSlice()
	if !reflect.DeepEqual(result, []int{2, 3, 1}) {
		t.Fatalf("Failed to invert: %v", result)
	}
}
