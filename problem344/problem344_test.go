package problem344

import (
	"reflect"
	"testing"
)

func TestCaseOne(t *testing.T) {
	input := []byte{'h', 'e', 'l', 'l', 'o'}
	output := []byte{'o', 'l', 'l', 'e', 'h'}

	reverseString(input)
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("Expected %v, but we have %v", input, output)
	}
}

func TestCaseTwo(t *testing.T) {
	input := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	output := []byte{'h', 'a', 'n', 'n', 'a', 'H'}

	reverseString(input)
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("Expected %v, but we have %v", input, output)
	}
}
