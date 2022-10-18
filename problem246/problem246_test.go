package problem246

import "testing"

func TestCase1(t *testing.T) {
	input := "69"
	if !isStrobogrammatic(input) {
		t.Errorf("%v is strobogrammatic", input)
	}
}

func TestCase2(t *testing.T) {
	input := "88"
	if !isStrobogrammatic(input) {
		t.Errorf("%v is strobogrammatic", input)
	}
}

func TestCase3(t *testing.T) {
	input := "962"
	if isStrobogrammatic(input) {
		t.Errorf("%v is strobogrammatic", input)
	}
}

func TestCase4(t *testing.T) {
	input := "6"
	if isStrobogrammatic(input) {
		t.Errorf("%v is strobogrammatic", input)
	}
}
