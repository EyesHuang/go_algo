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

	n := doublylinkedlist.NewNode(5, list.Head, nil)
	if data, err := list.Remove(n); err == nil {
		fmt.Println(data)
	} else {
		_ = fmt.Errorf(err.Error())
	}

	if !list.IsEmpty() {
		fmt.Println(list)
	}
}
