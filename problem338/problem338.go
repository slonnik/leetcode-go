package problem338

import (
	"fmt"
	"strings"
)

func countBits(n int) []int {
	result := []int{0}
	var index int
	for index = 1; index <= n; index++ {

		result = append(result, strings.Count(fmt.Sprintf("%b", index), "1"))
	}
	return result
}
