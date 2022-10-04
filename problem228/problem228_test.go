package problem228

import (
	"reflect"
	"testing"
)

func TestSummaryRangeTrivial(t *testing.T) {
	input := []int{0, 1, 2, 4, 5, 7}
	result := summaryRanges(input)
	expectedResult := []string{"0->2", "4->5", "7"}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Failed to test\n expected: %v\n but: %v", expectedResult, result)
	}
}

func TestSummaryNonTrivial(t *testing.T) {
	input := []int{0, 2, 3, 4, 6, 8, 9}
	result := summaryRanges(input)
	expectedResult := []string{"0", "2->4", "6", "8->9"}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Fatalf("Failed to test\n expected: %v\n but: %v", expectedResult, result)
	}
}
