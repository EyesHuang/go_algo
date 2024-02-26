package queue

import (
	"container/list"
	"errors"
)

type ListQueue[T any] struct {
	List *list.List
}

func NewListQueue[T any]() *ListQueue[T] {
	return &ListQueue[T]{
		list.New(),
	}
}

func (lq *ListQueue[T]) Size() int {
	return lq.List.Len()
}

func (lq *ListQueue[T]) IsEmpty() bool {
	return lq.Size() == 0
}

func (lq *ListQueue[T]) Offer(elem T) {
	lq.List.PushBack(elem)
}

func (lq *ListQueue[T]) Poll() (T, error) {
	var zero T
	if lq.IsEmpty() {
		return zero, errors.New("queue is empty")
	}

	elem := lq.List.Front()
	lq.List.Remove(elem)

	return elem.Value.(T), nil
}

func (lq *ListQueue[T]) Peek() (T, error) {
	var zero T
	if lq.IsEmpty() {
		return zero, errors.New("queue is empty")
	}

	return lq.List.Front().Value.(T), nil
}
