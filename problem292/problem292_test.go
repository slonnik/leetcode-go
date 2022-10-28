package problem292

import "testing"

func TestCaseOne(t *testing.T) {
	result := canWinNim(4)
	if result {
		t.Fatal("Expecting false")
	}
}

func TestCaseTwo(t *testing.T) {
	result := canWinNim(1)
	if !result {
		t.Fatal("Expecting true")
	}
}

func TestCaseThree(t *testing.T) {
	result := canWinNim(2)
	if !result {
		t.Fatal("Expecting true")
	}
}
