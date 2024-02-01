package doublylinkedlist

import (
	"errors"
	"reflect"
)

type Node[T any] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

type DoublyLinkedList[T any] struct {
	S    int
	Head *Node[T]
	Tail *Node[T]
}

func NewNode[T any](data T, prev, next *Node[T]) *Node[T] {
	return &Node[T]{
		data: data,
		prev: prev,
		next: next,
	}
}

func (list *DoublyLinkedList[T]) Size() int {
	return list.S
}

func (list *DoublyLinkedList[T]) IsEmpty() bool {
	return list.S == 0
}

func (list *DoublyLinkedList[T]) Clear() {
	t := list.Head

	for t != nil {
		next := t.next
		t.prev = nil
		t.next = nil
		t.data = *new(T)
		t = next
	}

	list.Head = nil
	list.Tail = nil
	list.S = 0

	/* Method 2:
	// In Go, when you set a pointer to nil, it effectively becomes unreachable and will be collected by the garbage collector
	// if there are no other references to it. This is a more idiomatic approach in Go and is more efficient.
	list.Head = nil
	list.Tail = nil
	list.S = 0
	*/
}

func (list *DoublyLinkedList[T]) AddLast(elem T) {
	if list.IsEmpty() {
		list.Head = NewNode(elem, nil, nil)
		list.Tail = list.Head
	} else {
		list.Tail.next = NewNode(elem, list.Tail, nil)
		list.Tail = list.Tail.next
	}
	list.S++
}

func (list *DoublyLinkedList[T]) AddFirst(elem T) {
	if list.IsEmpty() {
		list.Head = NewNode(elem, nil, nil)
		list.Tail = list.Head
	} else {
		list.Head.prev = NewNode(elem, nil, list.Head)
		list.Head = list.Head.prev
	}
	list.S++
}

func (list *DoublyLinkedList[T]) Add(elem T) {
	list.AddLast(elem)
}

func (list *DoublyLinkedList[T]) AddAt(index int, data T) {
	if index < 0 || index > list.S {
		return
	}

	switch index {
	case 0:
		list.AddFirst(data)
		break
	case list.S:
		list.AddLast(data)
		break
	}

	t := list.Head
	for i := 0; i < index; i++ {
		t = t.next
	}

	prev := t.prev
	n := NewNode(data, prev, t)

	prev.next = n
	t.prev = n

	list.S++
}

func (list *DoublyLinkedList[T]) PeekFirst() T {
	var zero T

	if list.IsEmpty() {
		return zero
	}
	return list.Head.data
}

func (list *DoublyLinkedList[T]) PeekLast() T {
	var zero T

	if list.IsEmpty() {
		return zero
	}
	return list.Tail.data
}

func (list *DoublyLinkedList[T]) RemoveFirst() (T, error) {
	var zero T
	if list.IsEmpty() {
		return zero, errors.New("empty list")
	}

	d := list.Head.data
	list.Head = list.Head.next
	list.S--

	if list.IsEmpty() {
		list.Tail = nil
	} else {
		list.Head.prev = nil
	}

	return d, nil
}

func (list *DoublyLinkedList[T]) RemoveLast() (T, error) {
	var zero T
	if list.IsEmpty() {
		return zero, errors.New("empty list")
	}

	d := list.Tail.data

	list.Tail = list.Tail.prev
	list.S--

	if list.IsEmpty() {
		list.Head = nil
	} else {
		list.Tail.next = nil
	}

	return d, nil
}

func (list *DoublyLinkedList[T]) RemoveByNode(node *Node[T]) (T, error) {
	var zero T
	if list.IsEmpty() {
		return zero, errors.New("empty list")
	}

	if node.prev == nil {
		return list.RemoveFirst()
	} else if node.next == nil {
		return list.RemoveLast()
	} else {
		prev := node.prev
		next := node.next

		prev.next = next
		next.prev = prev

		node.prev = nil
		node.next = nil
		list.S--
		return node.data, nil
	}
}

func (list *DoublyLinkedList[T]) RemoveAt(index int) (T, error) {
	var zero T
	if index >= list.S || index < 0 {
		return zero, errors.New("index out of bounds")
	}

	var current *Node[T]
	if index < list.S/2 {
		current = list.Head
		for i := 0; i != index; i++ {
			current = current.next
		}
	} else {
		current = list.Tail
		for i := list.S - 1; i != index; i-- {
			current = current.prev
		}
	}

	return list.RemoveByNode(current)
}

func (list *DoublyLinkedList[T]) RemoveByValue(obj T) (bool, error) {
	if list.IsEmpty() {
		return false, errors.New("empty list")
	}

	t := list.Head
	find := false
	for i := 0; i < list.S; i++ {
		if reflect.DeepEqual(t.data, obj) {
			find = true
			break
		}
		t = t.next
	}

	if !find {
		return false, errors.New("not found the object")
	}

	if t.prev == nil {
		if _, err := list.RemoveFirst(); err == nil {
			return true, nil
		} else {
			return false, err
		}
	} else if t.next == nil {
		if _, err := list.RemoveLast(); err == nil {
			return true, nil
		} else {
			return false, err
		}
	} else {
		prev := t.prev
		next := t.next

		prev.next = next
		next.prev = prev
		t.prev = nil
		t.next = nil
		list.S--
		return true, nil
	}
}
