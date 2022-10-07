package problem331

import "testing"

func TestTrivialCase(t *testing.T) {
	if !isPowerOfTwo(1) {
		t.Fatalf("1 is power of two")
	}
}

func TestTrivialCase2(t *testing.T) {
	if !isPowerOfTwo(8) {
		t.Fatalf("8 is power of two")
	}
}

func TestTrivialCase3(t *testing.T) {
	if isPowerOfTwo(3) {
		t.Fatalf("3 is not power of two")
	}
}
