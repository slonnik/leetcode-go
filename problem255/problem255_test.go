package problem255

import (
	"testing"
)

func TestProblem_1(t *testing.T) {

	var stack = Constructor()
	stack.Push(1)
	stack.Push(2)
	if stack.Top() != 2 {
		t.Fatalf("Invalid 'Top' operation")
	}

	if stack.Pop() != 2 {
		t.Fatalf("Invalid 'Pop' operation")
	}

	if stack.Empty() {
		t.Fatalf("Invalid 'Empty' operation")
	}
}
