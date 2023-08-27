package stack

import (
	"testing"
)

func TestArrayStack_PushPop(t *testing.T) {
	stack := NewArrayStack()

	_ = stack.Push(1)
	_ = stack.Push(2)
	_ = stack.Push(3)

	if top, _ := stack.Pop(); top != 3 {
		t.Errorf("Expected top value to be 3, got %v", top)
	}

	if top, _ := stack.Pop(); top != 2 {
		t.Errorf("Expected top value to be 2, got %v", top)
	}

	if top, _ := stack.Pop(); top != 1 {
		t.Errorf("Expected top value to be 1, got %v", top)
	}

	if _, err := stack.Pop(); err == nil {
		t.Errorf("Expected error when popping from empty stack, but got nil")
	}
}

func TestArrayStack_Peek(t *testing.T) {
	stack := NewArrayStack()

	_ = stack.Push(5)

	if top, _ := stack.Peek(); top != 5 {
		t.Errorf("Expected top value to be 5, got %v", top)
	}

	if top, _ := stack.Pop(); top != 5 {
		t.Errorf("Expected top value to be 5, got %v", top)
	}

	if _, err := stack.Peek(); err == nil {
		t.Errorf("Expected error when peeking into empty stack, but got nil")
	}
}

func TestArrayStack_Empty(t *testing.T) {
	stack := NewArrayStack()

	if !stack.Empty() {
		t.Error("Expected stack to be empty, but it's not")
	}

	_ = stack.Push(10)

	if stack.Empty() {
		t.Error("Expected stack not to be empty, but it is")
	}
}
