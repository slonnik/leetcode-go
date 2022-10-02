package problem226

import (
	"github.com/slonnik/leetcode-go/util"
	"testing"
)

func TestInvertTrivial(t *testing.T) {
	node := util.TreeNodeFromSlice([]int{2, 1, 3})
	InvertTree(node)
}
