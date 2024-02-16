package doublylinkedlist

import (
	"testing"
)

func TestDoublyLinkedList_AddFirst(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.AddFirst(1)

	if list.Size() != 1 {
		t.Errorf("Expected list size of 1, got %d", list.Size())
	}

	if val, _ := list.PeekFirst(); val != 1 {
		t.Errorf("Expected first element to be 1, got %d", val)
	}

	list.AddFirst(2)
	if val, _ := list.PeekFirst(); val != 2 {
		t.Errorf("Expected first element to be 2, got %d", val)
	}
}

func TestDoublyLinkedList_AddLast(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.AddLast(1)

	if list.Size() != 1 {
		t.Errorf("Expected list size of 1, got %d", list.Size())
	}

	if val, _ := list.PeekLast(); val != 1 {
		t.Errorf("Expected last element to be 1, got %d", val)
	}

	list.AddLast(2)
	if val, _ := list.PeekLast(); val != 2 {
		t.Errorf("Expected last element to be 2, got %d", val)
	}
}

func TestDoublyLinkedList_Clear(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.AddLast(1)
	list.AddLast(2)
	list.Clear()

	if !list.IsEmpty() {
		t.Error("Expected list to be empty after clear")
	}

	if list.Head != nil || list.Tail != nil {
		t.Error("Expected Head and Tail to be nil after clear")
	}
}

func TestDoublyLinkedList_Size(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}

	list.AddLast(1)
	list.AddLast(2)

	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
}

func TestDoublyLinkedList_AddAt(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	list.Add(1)
	list.Add(2)

	if err := list.AddAt(2, 3); err != nil || list.Size() != 3 {
		t.Errorf("Expected to add element at the end, got size %d, error: %v", list.Size(), err)
	}
}

func TestDoublyLinkedList_RemoveFirst(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	if _, err := list.RemoveFirst(); err == nil {
		t.Error("Expected an error when removing from an empty list")
	}

	list.AddFirst(1)
	list.AddFirst(2)

	if val, err := list.RemoveFirst(); err != nil || val != 2 {
		t.Errorf("Expected to remove 2, got %d", val)
	}

	if list.Size() != 1 {
		t.Errorf("Expected list size of 1, got %d", list.Size())
	}
}

func TestDoublyLinkedList_RemoveLast(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	if _, err := list.RemoveLast(); err == nil {
		t.Errorf("Expected to have an error when removing from an empty list")
	}

	list.Add(1)
	list.Add(2)

	if val, err := list.RemoveLast(); err != nil || val != 2 {
		t.Errorf("Expected to remove element at the end, got %d, error: %v", val, err)
	}
}

func TestDoublyLinkedList_RemoveByNode(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	newNode := NewNode(1, nil, nil)

	if _, err := list.RemoveByNode(newNode); err == nil {
		t.Errorf("Expected to have an error when removing from an empty list")
	}

	list.Add(1)
	list.Add(2)
	newNode = list.Tail

	if val, err := list.RemoveByNode(newNode); err != nil || val != 2 {
		t.Errorf("Expected to remove value 2, got %d, error: %v", val, err)
	}
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	if _, err := list.RemoveAt(0); err == nil {
		t.Errorf("Expected to have an error when removing an empty list")
	}

	list.Add(1)
	list.Add(2)

	if val, err := list.RemoveAt(1); err != nil || val != 2 {
		t.Errorf("Expected to remove element at index 1, got %d, error = %v", val, err)
	}
}

func TestDoublyLinkedList_RemoveByValue(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	if _, err := list.RemoveByValue(1); err == nil {
		t.Errorf("Expected to have an error when removing from an empty list")
	}

	for i := 0; i < 5; i++ {
		list.Add(i)
	}

	if found, err := list.RemoveByValue(3); err != nil || found != true {
		t.Errorf("Expected to remove value 3, not found or error: %v", err)
	}
}

func TestDoublyLinkedList_IndexOf(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	for i := 0; i < 5; i++ {
		list.Add(i)
	}

	if at := list.IndexOf(3); at != 3 {
		t.Errorf("Expected to get value 3 at index 3, got %d", at)
	}
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)

	if !list.Contains(2) {
		t.Errorf("Expected list to contain 2")
	}

	if list.Contains(4) {
		t.Errorf("Did not expect list to contain 4")
	}
}

func TestDoublyLinkedList_PeekFirst(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	if _, err := list.PeekFirst(); err == nil {
		t.Errorf("Expected error when peeking at an empty list")
	}

	list.AddLast(1)
	list.AddLast(2)

	if val, err := list.PeekFirst(); err != nil || val != 1 {
		t.Errorf("Expected first element to be 1, got %d, error: %v", val, err)
	}
}

func TestDoublyLinkedList_PeekLast(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	if _, err := list.PeekLast(); err == nil {
		t.Errorf("Expected error when peeking at an empty list")
	}

	list.AddLast(1)
	list.AddLast(2)

	if val, err := list.PeekLast(); err != nil || val != 2 {
		t.Errorf("Expected last element to be 2, got %d, error: %v", val, err)
	}
}

func TestDoublyLinkedList_IsEmpty(t *testing.T) {
	list := &DoublyLinkedList[int]{}

	if !list.IsEmpty() {
		t.Errorf("Expected new list to be empty")
	}

	list.AddLast(1)

	if list.IsEmpty() {
		t.Errorf("Did not expect list with elements to be empty")
	}

	if val, err := list.RemoveFirst(); err != nil || val != 1 {
		t.Errorf("Expected to remove the last element, got %d, error: %v", val, err)
	}

	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty after removing the only element")
	}
}
