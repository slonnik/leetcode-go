package problem349

import (
	"reflect"
	"testing"
)

func TestCaseOne(t *testing.T) {

	result := intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})

	if !reflect.DeepEqual(result, []int{4, 9}) {
		t.Fatalf("invalid intersection, expected {4,9}, but have %v", result)
	}
}
