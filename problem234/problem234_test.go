package problem234

import (
	"github.com/slonnik/leetcode-go/util"
	"testing"
)

func TestTrivialPalindrome(t *testing.T) {
	data := []int{1, 2, 2, 1}
	node := util.ListNodeFromSlice(data)
	if !isPalindrome(node) {
		t.Fatalf("%v is palindrome", data)
	}
}

func TestTrivialPalindrome2(t *testing.T) {
	data := []int{1, 2, 3, 1}
	node := util.ListNodeFromSlice(data)
	if isPalindrome(node) {
		t.Fatalf("%v is no palindrome", data)
	}
}
