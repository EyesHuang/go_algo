package bst

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Value       T
	Left, Right *Node[T]
}

type BinarySearchTree[T constraints.Ordered] struct {
	Root *Node[T]
}

func NewNode[T constraints.Ordered](value T) *Node[T] {
	return &Node[T]{Value: value}
}

func NewBinarySearchTree[T constraints.Ordered](value ...T) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

// Insert adds a new Node with the specified value into the tree.
func (bst *BinarySearchTree[T]) Insert(value T) {
	newNode := NewNode(value)

	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	currentNode := bst.Root

	for {
		if value > currentNode.Value {
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return
			}
			currentNode = currentNode.Right
		} else {
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return
			}
			currentNode = currentNode.Left
		}
	}
}

// LookUp searches for a node with the given value and returns it if found.
func (bst *BinarySearchTree[T]) LookUp(value T) *Node[T] {
	currentNode := bst.Root

	for currentNode != nil {
		if value > currentNode.Value {
			currentNode = currentNode.Right
		} else if value < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			return currentNode
		}
	}

	return nil
}
