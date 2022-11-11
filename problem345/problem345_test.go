package problem345

import (
	"testing"
)

func TestCaseOne(t *testing.T) {

	input := "hello"
	expected := "holle"
	result := reverseVowels(input)

	if result != expected {
		t.Fatalf("Expected %v, but result is %v", expected, result)
	}
}

func TestCaseTwo(t *testing.T) {

	input := "leetcode"
	expected := "leotcede"
	result := reverseVowels(input)

	if result != expected {
		t.Fatalf("Expected %v, but result is %v", expected, result)
	}
}
