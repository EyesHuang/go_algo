package queue

import "testing"

func TestListQueue_Size(t *testing.T) {
	q := NewListQueue[int]()

	if !q.IsEmpty() {
		t.Errorf("Expected size 0, got %d", q.Size())
	}

	q.Offer(1)
	q.Offer(2)

	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}
}

func TestListQueue_IsEmpty(t *testing.T) {
	q := NewListQueue[int]()
	q.Offer(1)
	q.Offer(2)

	if q.IsEmpty() {
		t.Errorf("Expected queue not empty")
	}

	if elem, err := q.Poll(); err != nil {
		t.Errorf("Expected the element to be 1, got %d, error %v", elem, err)
	}

	if q.IsEmpty() {
		t.Errorf("Expected queue not empty")
	}
}

func TestListQueue_Offer(t *testing.T) {
	q := NewListQueue[int]()
	q.Offer(1)
	q.Offer(2)

	if q.Size() != 2 {
		t.Errorf("Expected queue size of 2, got %d", q.Size())
	}

	if elem, err := q.Peek(); err != nil || elem != 1 {
		t.Errorf("Expected the element to be 1, got %d", elem)
	}
}

func TestListQueue_Poll(t *testing.T) {
	q := NewListQueue[int]()

	if elem, err := q.Poll(); err == nil || elem != 0 {
		t.Errorf("Expected empty queue to have error when polling")
	}

	q.Offer(1)
	q.Offer(2)

	if elem, err := q.Poll(); err != nil || elem != 1 {
		t.Errorf("Expected the element to be 1, got %d", elem)
	}
}

func TestListQueue_Peek(t *testing.T) {
	q := NewListQueue[int]()

	if elem, err := q.Peek(); err == nil || elem != 0 {
		t.Errorf("Expected emptyt queue to have error when peeking")
	}

	q.Offer(1)

	if elem, err := q.Peek(); err != nil || elem != 1 {
		t.Errorf("Expected the element to be 1, got %d, error = %v", elem, err)
	}
}
