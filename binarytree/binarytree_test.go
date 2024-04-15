package binarytree

import (
	"fmt"
	"testing"
)

const data = "A B C D E F x x x G H x I"

// D B G E H A F I C
func TestBinaryTree_InorderByParent(t *testing.T) {
	tree := NewBinaryTree(data)
	tree.InorderByParent()
	fmt.Println()
}

// L D M B G E H A N F I C K
func TestBinaryTree_InsertLevelOrder(t *testing.T) {
	tree := NewBinaryTree(data)

	tree.InsertLevelOrder('K')
	tree.InsertLevelOrder('L')
	tree.InsertLevelOrder('M')
	tree.InsertLevelOrder('N')

	tree.InorderByParent()
	fmt.Println()
}

// D B G E H A F I C
func TestBinaryTree_InOrder(t *testing.T) {
	tree := NewBinaryTree(data)
	tree.InOrder()
	fmt.Println()
}

// A B D E G H C F I
func TestBinaryTree_PreOrder(t *testing.T) {
	tree := NewBinaryTree(data)
	tree.PreOrder()
	fmt.Println()
}

// D G H E B I F C A
func TestBinaryTree_PostOrder(t *testing.T) {
	tree := NewBinaryTree(data)
	tree.PostOrder()
	fmt.Println()
}
