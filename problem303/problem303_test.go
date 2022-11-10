package problem303

import (
	"testing"
)

func TestMainCase(t *testing.T) {
	data := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(data)
	if numArray.SumRange(0, 2) != 1 {
		t.Fatalf("Expected 1")
	}

	if numArray.SumRange(2, 5) != -1 {
		t.Fatalf("Expected -1")
	}

	if numArray.SumRange(0, 5) != -3 {
		t.Fatalf("Expected -3")
	}

	if numArray.SumRange(2, 2) != 3 {
		t.Fatalf("Expected 2")
	}
}
