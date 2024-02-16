package stack

type Stack[T any] interface {
	Size() int
	IsEmpty() bool
	Push(elem T)
	Pop() (T, error)
	Peek() (T, error)
}
