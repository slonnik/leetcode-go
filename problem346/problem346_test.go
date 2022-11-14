package problem346

import (
	"testing"
)

func TestCaseOne(t *testing.T) {
	obj := Constructor(3)
	if obj.Next(1) != 1 {
		t.Errorf("Expecting 1")
	}

	if obj.Next(10) != 5.5 {
		t.Errorf("Expecting 5.5")
	}

	if obj.Next(3) != 4.66667 {
		t.Errorf("Expecting 4.66667")
	}

	if obj.Next(5) != 6 {
		t.Errorf("Expecting 6")
	}
}
