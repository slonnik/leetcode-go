package problem242

import "testing"

func TestTrivialCase1(t *testing.T) {
	if !isAnagram("anagram", "nagaram") {
		t.Fatalf("Words are not nagarams")
	}
}

func TestTrivialCase2(t *testing.T) {
	if isAnagram("rat", "car") {
		t.Fatalf("Words are nagarams")
	}
}
