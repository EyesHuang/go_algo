package binarysearchtree

import (
	"errors"
	"fmt"
)

// TreeNode structure
type TreeNode struct {
	key        int
	element    string
	leftChild  *TreeNode
	rightChild *TreeNode
	parent     *TreeNode
}

// NewTreeNode creates a new TreeNode
func NewTreeNode(key int, element string) *TreeNode {
	return &TreeNode{
		key:     key,
		element: element,
	}
}

// BST structure
type BST struct {
	root *TreeNode
}

// NewBST creates a new BST
func NewBST() *BST {
	return &BST{}
}

// InsertBST inserts a new node into the BST
func (bst *BST) InsertBST(key int, element string) {
	newNode := NewTreeNode(key, element)

	var y *TreeNode
	x := bst.root

	for x != nil {
		y = x

		if newNode.key > x.key {
			x = x.rightChild
		} else if newNode.key < x.key {
			x = x.leftChild
		}
	}

	newNode.parent = y

	if y == nil {
		bst.root = newNode
	} else if newNode.key > y.key {
		y.rightChild = newNode
	} else if newNode.key < y.key {
		y.leftChild = newNode
	}
}

// Search finds a node by its key in the BST
func (bst *BST) Search(key int) *TreeNode {
	if bst.root == nil {
		return nil
	}

	current := bst.root

	for current != nil {
		if current.key == key {
			return current
		} else if key < current.key {
			current = current.leftChild
		} else if key > current.key {
			current = current.rightChild
		}
	}

	return nil
}

func (bst *BST) InorderPrint() {
	current := leftMost(bst.root)

	for current != nil {
		fmt.Printf("%s(%d) ", current.element, current.key)
		current = inorderSuccessor(current)
	}
}

func leftMost(node *TreeNode) *TreeNode {
	for node != nil && node.leftChild != nil {
		node = node.leftChild
	}
	return node
}

func inorderSuccessor(node *TreeNode) *TreeNode {
	if node.rightChild != nil {
		return leftMost(node.rightChild)
	}

	successor := node.parent

	for successor != nil && node == successor.rightChild {
		node = successor
		successor = successor.parent
	}

	return successor
}

func (bst *BST) LevelOrder() {
	queue := []*TreeNode{bst.root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Printf("%s(%d) ", current.element, current.key)

		if current.leftChild != nil {
			queue = append(queue, current.leftChild)
		}

		if current.rightChild != nil {
			queue = append(queue, current.rightChild)
		}
	}
}

// DeleteBST deletes a node with the specified key in the BST
func (bst *BST) DeleteBST(key int) error {
	deleteNode := bst.Search(key) // Use the search function to find the node
	if deleteNode == nil {
		return errors.New("data not found")
	}

	y, x := bst.determineNodeToDelete(deleteNode)

	// Reconnect the child x with the rest of the tree
	bst.reconnectChild(x, y)

	// If y is not the delete node, swap the contents
	if y != deleteNode {
		bst.swapContents(y, deleteNode)
	}

	return nil
}

// Helper function to determine the node to delete
func (bst *BST) determineNodeToDelete(node *TreeNode) (*TreeNode, *TreeNode) {
	if node.leftChild == nil || node.rightChild == nil {
		return node, bst.getChild(node)
	} else {
		successor := inorderSuccessor(node)
		return successor, bst.getChild(successor)
	}
}

// Helper function to get the child of a node
func (bst *BST) getChild(node *TreeNode) *TreeNode {
	if node.leftChild != nil {
		return node.leftChild
	}
	return node.rightChild
}

// Helper function to reconnect child with the BST
func (bst *BST) reconnectChild(child, node *TreeNode) {
	// Reconnect the child x with the rest of the tree
	if child != nil {
		child.parent = node.parent
	}

	// If y is the root, update the root to x
	if node.parent == nil {
		bst.root = child
	} else if node == node.parent.leftChild {
		node.parent.leftChild = child
	} else {
		node.parent.rightChild = child
	}
}

// Helper function to swap the contents of two nodes
func (bst *BST) swapContents(src, dest *TreeNode) {
	dest.key = src.key
	dest.element = src.element
}
