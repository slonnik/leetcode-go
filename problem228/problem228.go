package problem228

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/summary-ranges/

func summaryRanges(nums []int) []string {

	var result []string

	if len(nums) == 0 {
		return result
	}

	startIndex := 0

	for index, value := range nums {
		if index == 0 {
			continue
		}
		if math.Abs(float64(value-nums[index-1])) > 1 {
			if startIndex == index-1 {
				result = append(result, fmt.Sprint(nums[startIndex]))
			} else {
				result = append(result, fmt.Sprintf("%v->%v", nums[startIndex], nums[index-1]))
			}
			startIndex = index
		}
	}

	if startIndex == len(nums)-1 {
		result = append(result, fmt.Sprint(nums[startIndex]))
	} else {
		result = append(result, fmt.Sprintf("%v->%v", nums[startIndex], nums[len(nums)-1]))
	}

	return result
}
