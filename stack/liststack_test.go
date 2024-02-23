package stack

import "testing"

func TestListStack_Size(t *testing.T) {
	stack := NewListStack[int]()

	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}

	stack.Push(1)
	stack.Push(2)

	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}
}

func TestListStack_IsEmpty(t *testing.T) {
	stack := NewListStack[int]()

	if !stack.IsEmpty() {
		t.Errorf("Expected empty stack")
	}

	stack.Push(1)
	stack.Push(2)

	if stack.IsEmpty() {
		t.Errorf("Expected stack not empty")
	}

	if elem, err := stack.Pop(); err != nil || elem != 2 {
		t.Errorf("Expected the element to be 2, got %d, error %v", elem, err)
	}

	if stack.IsEmpty() {
		t.Errorf("Expected stack not empty")
	}
}

func TestListStack_Push(t *testing.T) {
	stack := NewListStack[int]()
	stack.Push(1)
	stack.Push(2)

	if stack.Size() != 2 {
		t.Errorf("Expected stack size of 2, got %d", stack.Size())
	}

	if elem, err := stack.Peek(); err != nil || elem != 2 {
		t.Errorf("Expected the element to be 2, got %d", elem)
	}
}

func TestListStack_Pop(t *testing.T) {
	stack := NewListStack[int]()

	if elem, err := stack.Pop(); err == nil || elem != 0 {
		t.Errorf("Expected empty stack to have error when popping up")
	}

	stack.Push(1)

	if elem, err := stack.Pop(); err != nil || elem != 1 {
		t.Errorf("Expected the element to be 1, got %d, error %v", elem, err)
	}
}

func TestListStack_Peek(t *testing.T) {
	stack := NewListStack[int]()

	if elem, err := stack.Peek(); err == nil || elem != 0 {
		t.Errorf("Expected empty stack to have error")
	}

	stack.Push(1)

	if elem, err := stack.Peek(); err != nil || elem != 1 {
		t.Errorf("Expected the element to be 1, got %d, err %v", elem, err)
	}
}
