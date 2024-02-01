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

	index := list.IndexOf(10)
	fmt.Println(index)

	if !list.IsEmpty() {
		fmt.Println(list)
	}
}
