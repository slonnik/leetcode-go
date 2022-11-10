package problem326

import (
	"testing"
)

func TestCaseOne(t *testing.T) {
	if !isPowerOfThree(27) {
		t.Fatal("Must be true")
	}

}

func TestCaseTwo(t *testing.T) {

	if isPowerOfThree(-1) {
		t.Fatal("Must be false")
	}
}

func TestCaseThree(t *testing.T) {

	if isPowerOfThree(0) {
		t.Fatal("Must be false")
	}
}
