package stack

import (
	"container/list"
	"errors"
	"reflect"
)

type ListStack[T any] struct {
	list *list.List
}

func NewListStack[T any]() *ListStack[T] {
	return &ListStack[T]{
		list: list.New(),
	}
}

func (ls *ListStack[T]) Size() int {
	return ls.list.Len()
}

func (ls *ListStack[T]) IsEmpty() bool {
	return ls.Size() == 0
}

func (ls *ListStack[T]) Push(elem T) {
	ls.list.PushBack(elem)
}

func (ls *ListStack[T]) Pop() (T, error) {
	var zero T
	if ls.IsEmpty() {
		return zero, errors.New("stack is empty")
	}

	elem := ls.list.Back()
	v := ls.list.Remove(elem)
	if !reflect.DeepEqual(v, elem.Value) {
		return zero, errors.New("error in removing the last element")
	}

	return elem.Value.(T), nil
}

func (ls *ListStack[T]) Peek() (T, error) {
	var zero T
	if ls.IsEmpty() {
		return zero, errors.New("stack is empty")
	}

	elem := ls.list.Back()
	return elem.Value.(T), nil
}
