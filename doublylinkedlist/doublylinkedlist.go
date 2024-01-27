package doublylinkedlist

import "errors"

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

func (list *DoublyLinkedList[T]) RemoveFirst() T {
	var zero T
	if list.IsEmpty() {
		return zero
	}

	d := list.Head.data
	list.Head = list.Head.next
	list.S--

	if list.IsEmpty() {
		list.Tail = nil
	} else {
		list.Head.prev = nil
	}

	return d
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
