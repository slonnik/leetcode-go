package problem283

import (
	"reflect"
	"testing"
)

func TestCaseOne(t *testing.T) {
	data := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but we have %v", expected, data)
	}
}

func TestCaseTwo(t *testing.T) {
	data := []int{0}
	expected := []int{0}
	moveZeroes(data)

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, but we have %v", expected, data)
	}
}
