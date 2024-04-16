package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestBST_InorderPrint(t *testing.T) {
	bst := genBST()
	bst.InorderPrint()
	fmt.Println()
}

func TestBST_LevelOrder(t *testing.T) {
	bst := genBST()
	bst.LevelOrder()
	fmt.Println()
}

func TestBST_Search_Success(t *testing.T) {
	bst := genBST()
	node := bst.Search(1000)
	if node != nil {
		fmt.Printf("There is %s(%d)\n", node.element, node.key)
	} else {
		fmt.Println("No element with Key(1000)")
	}
}

func TestBST_Search_Fail(t *testing.T) {
	bst := genBST()
	node := bst.Search(73)
	if node != nil {
		fmt.Printf("There is %s(%d)\n", node.element, node.key)
	} else {
		fmt.Println("No element with Key(73)")
	}
}

func genBST() *BST {
	bst := NewBST()
	bst.InsertBST(8, "Master Roshi")
	bst.InsertBST(1000, "Goku")
	bst.InsertBST(2, "Krillin")
	bst.InsertBST(513, "Piccolo")

	return bst
}
