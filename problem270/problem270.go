package problem270

import (
	"github.com/slonnik/leetcode-go/util"
	"math"
)

func closestValue(root *util.TreeNode, target float64) int {

	result := root.Val
	internalClosestValue(root, target, &result)
	return result
}

func internalClosestValue(node *util.TreeNode, target float64, minValue *int) {
	if node == nil {
		return
	}

	minDelta := math.Abs(float64(*minValue) - target)
	currentDelta := math.Abs(float64(node.Val) - target)

	if currentDelta < minDelta {
		*minValue = node.Val
	}
	internalClosestValue(node.Left, target, minValue)
	internalClosestValue(node.Right, target, minValue)
}
