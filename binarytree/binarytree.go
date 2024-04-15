package binarytree

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	leftChild  *TreeNode
	rightChild *TreeNode
	parent     *TreeNode
	data       rune
}

type BinaryTree struct {
	root *TreeNode
}

func NewTreeNode(data rune) *TreeNode {
	return &TreeNode{data: data}
}

func NewBinaryTree(data string) *BinaryTree {
	tree := &BinaryTree{}
	tree.LevelOrderConstruct(data)
	return tree
}

func (bt *BinaryTree) LevelOrderConstruct(data string) {
	fields := strings.Fields(data)
	if len(fields) == 0 {
		return
	}

	rootData := rune(fields[0][0])
	if rootData == 'x' { // Assuming 'x' stands for no node
		return
	}

	bt.root = NewTreeNode(rootData)
	queue := []*TreeNode{bt.root}
	index := 1

	for len(queue) > 0 && index < len(fields) {
		current := queue[0]
		queue = queue[1:]

		// Handle left child
		if index < len(fields) && fields[index][0] != 'x' {
			current.leftChild = NewTreeNode(rune(fields[index][0]))
			current.leftChild.parent = current
			queue = append(queue, current.leftChild)
		}
		index++

		// Handle right child
		if index < len(fields) && fields[index][0] != 'x' {
			current.rightChild = NewTreeNode(rune(fields[index][0]))
			current.rightChild.parent = current
			queue = append(queue, current.rightChild)
		}
		index++
	}
}

func (bt *BinaryTree) InsertLevelOrder(data rune) {
	newNode := NewTreeNode(data)
	queue := []*TreeNode{bt.root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.leftChild == nil {
			current.leftChild = newNode
			newNode.parent = current
			return
		} else {
			queue = append(queue, current.leftChild)
		}

		if current.rightChild == nil {
			current.rightChild = newNode
			newNode.parent = current
			return
		} else {
			queue = append(queue, current.rightChild)
		}
	}
}

func (bt *BinaryTree) PreOrder() {
	preOrder(bt.root)
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%c ", node.data)
	preOrder(node.leftChild)
	preOrder(node.rightChild)
}

func (bt *BinaryTree) InOrder() {
	inOrder(bt.root)
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.leftChild)
	fmt.Printf("%c ", node.data)
	inOrder(node.rightChild)
}

func (bt *BinaryTree) PostOrder() {
	postOrder(bt.root)
}

func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.leftChild)
	postOrder(node.rightChild)
	fmt.Printf("%c ", node.data)
}

func (bt *BinaryTree) LevelOrder() {
	if bt.root == nil {
		return
	}

	queue := []*TreeNode{bt.root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Printf("%c ", current.data)

		if current.leftChild != nil {
			queue = append(queue, current.leftChild)
		}

		if current.rightChild != nil {
			queue = append(queue, current.rightChild)
		}
	}
}

func leftMost(current *TreeNode) *TreeNode {
	for current.leftChild != nil {
		current = current.leftChild
	}
	return current
}

func inorderSuccessor(current *TreeNode) *TreeNode {
	if current.rightChild != nil {
		return leftMost(current.rightChild)
	}

	successor := current.parent

	for successor != nil && current == successor.rightChild {
		current = successor
		successor = successor.parent
	}

	return successor
}

func (bt *BinaryTree) InorderByParent() {
	current := leftMost(bt.root)

	for current != nil {
		fmt.Printf("%c ", current.data)
		current = inorderSuccessor(current)
	}
}

func rightMost(current *TreeNode) *TreeNode {
	for current.rightChild != nil {
		current = current.rightChild
	}
	return current
}

func inorderPredecessor(current *TreeNode) *TreeNode {
	if current.leftChild != nil {
		return rightMost(current.leftChild)
	}

	predecessor := current.parent

	for predecessor != nil && current == predecessor.leftChild {
		current = predecessor
		predecessor = predecessor.parent
	}

	return predecessor
}

func (bt *BinaryTree) InOrderReverse(root *TreeNode) {
	current := rightMost(root)

	for current != nil {
		fmt.Printf("%c ", current.data)
		current = inorderPredecessor(current)
	}
}
