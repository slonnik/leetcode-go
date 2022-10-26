package problem268

import (
	"testing"
)

func TestCaseOne(t *testing.T) {
	if 2 != missingNumber([]int{3, 0, 1}) {
		t.Fatalf("2 is missing")
	}
}

func TestCaseTwo(t *testing.T) {
	if 2 != missingNumber([]int{0, 1}) {
		t.Fatalf("2 is missing")
	}
}

func TestCaseThree(t *testing.T) {
	if 8 != missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) {
		t.Fatalf("8 is missing")
	}
}

func TestCaseFour(t *testing.T) {
	if 0 != missingNumber([]int{1}) {
		t.Fatalf("0 is missing")
	}
}
