package stack

import "errors"

type ArrayStack[T any] struct {
	data []T
}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{}
}

func (ar *ArrayStack[T]) Size() int {
	return len(ar.data)
}

func (ar *ArrayStack[T]) IsEmpty() bool {
	return len(ar.data) == 0
}

func (ar *ArrayStack[T]) Push(value T) {
	ar.data = append(ar.data, value)
}

func (ar *ArrayStack[T]) Pop() (T, error) {
	var zero T
	if ar.IsEmpty() {
		return zero, errors.New("stack is empty")
	}

	lastIdx := len(ar.data) - 1
	elem := ar.data[lastIdx]
	ar.data = ar.data[:lastIdx]
	return elem, nil
}

func (ar *ArrayStack[T]) Peek() (T, error) {
	var zero T
	if ar.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return ar.data[len(ar.data)-1], nil
}
