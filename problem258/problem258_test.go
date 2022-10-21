package problem258

import (
	"testing"
)

func TestCaseOne(t *testing.T) {
	expected := 2
	result := addDigits(38)
	if result != expected {
		t.Errorf("Expected %v, but result is %v", expected, result)
	}
}

func TestCaseTwo(t *testing.T) {
	expected := 0
	result := addDigits(0)
	if result != expected {
		t.Errorf("Expected %v, but result is %v", expected, result)
	}
}

func TestCaseThree(t *testing.T) {
	expected := 1
	result := addDigits(10)
	if result != expected {
		t.Errorf("Expected %v, but result is %v", expected, result)
	}
}
