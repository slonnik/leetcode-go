package problem338

import (
	"reflect"
	"testing"
)

func TestCaseOne(t *testing.T) {

	expectedData := []int{0, 1, 1, 2, 1, 2}
	result := countBits(5)
	if !reflect.DeepEqual(result, expectedData) {
		t.Fatalf("Expected %v, but was %v", expectedData, result)
	}
}
