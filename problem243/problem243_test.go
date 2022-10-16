package problem243

import "testing"

func TestCase1(t *testing.T) {
	distance := shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"},
		"coding", "practice")
	if distance != 3 {
		t.Fatalf("invalid distance %v", distance)
	}
}

func TestCase2(t *testing.T) {
	distance := shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"},
		"makes", "coding")
	if distance != 1 {
		t.Fatalf("invalid distance %v", distance)
	}
}
