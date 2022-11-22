package problem350

import (
	"reflect"
	"testing"
)

func TestCaseOne(t *testing.T) {

	result := intersection([]int{1, 2, 2, 1}, []int{2, 2})

	if !reflect.DeepEqual(result, []int{2, 2}) {
		t.Fatalf("invalid intersection, expected {2,2}, but have %v", result)
	}
}

func TestCaseTwo(t *testing.T) {

	result := intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	expected := []int{4, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("invalid intersection, expected %v, but have %v", expected, result)
	}
}
