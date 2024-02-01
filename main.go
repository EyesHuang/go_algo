package main

import (
	"algo/doublylinkedlist"
	"fmt"
)

func main() {
	h := doublylinkedlist.NewNode(4, nil, nil)
	list := doublylinkedlist.DoublyLinkedList[int]{
		S:    1,
		Head: h,
		Tail: h,
	}

	list.Add(5)
	list.Add(6)

	if check, err := list.RemoveByValue(5); err == nil {
		fmt.Println(check)
	}

	if !list.IsEmpty() {
		fmt.Println(list)
	}
}
