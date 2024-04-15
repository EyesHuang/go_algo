package bst

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test_BinarySearchTree_Main(t *testing.T) {
	bst := NewBinarySearchTree[int]()
	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	b, err := json.Marshal(bst)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	fmt.Println(bst.LookUp(20))
}
