package binarysearchtree

import "fmt"

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
