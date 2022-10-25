package problem266

import "testing"

func TestCaseOne(t *testing.T) {
	if canPermutePalindrome("code") {
		t.Errorf("Can't be palindrome")
	}
}

func TestCaseTwo(t *testing.T) {
	if !canPermutePalindrome("aab") {
		t.Errorf("Must be palindrome")
	}
}

func TestCaseThree(t *testing.T) {
	if !canPermutePalindrome("carerac") {
		t.Errorf("Must be palindrome")
	}
}
