package queue

import "errors"

type ArrayQueue[T any] struct {
	data []T
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{}
}

func (aq *ArrayQueue[T]) Size() int {
	return len(aq.data)
}

func (aq *ArrayQueue[T]) IsEmpty() bool {
	return aq.Size() == 0
}

func (aq *ArrayQueue[T]) Offer(elem T) {
	aq.data = append(aq.data, elem)
}

func (aq *ArrayQueue[T]) Poll() (T, error) {
	var zero T
	if aq.IsEmpty() {
		return zero, errors.New("queue is empty")
	}

	elem := aq.data[0]
	aq.data = aq.data[1:]
	return elem, nil
}

func (aq *ArrayQueue[T]) Peek() (T, error) {
	var zero T
	if aq.IsEmpty() {
		return zero, errors.New("queue is empty")
	}

	return aq.data[0], nil
}
