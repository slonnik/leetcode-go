package problem263

import (
	"testing"
)

func TestCaseOne(t *testing.T) {
	if !isUgly(1) {
		t.Error("1 is ugly")
	}
}

func TestCaseTwo(t *testing.T) {
	if !isUgly(6) {
		t.Error("6 is ugly")
	}
}

func TestCaseThree(t *testing.T) {
	if isUgly(14) {
		t.Error("14 is not ugly")
	}
}
