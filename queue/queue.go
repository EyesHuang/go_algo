package queue

type Queue[T any] interface {
	Size() int
	IsEmpty() bool
	Offer(elem T)
	Poll() (T, error)
	Peek() (T, error)
}
