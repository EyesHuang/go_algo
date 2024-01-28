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

	if data, err := list.RemoveAt(2); err == nil {
		fmt.Println(data)
	} else {
		fmt.Println(err.Error())
	}

	if !list.IsEmpty() {
		fmt.Println(list)
	}
}
