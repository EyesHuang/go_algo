package bst

import (
	"fmt"
)

type TreeNode struct {
	leftChild  *TreeNode
	rightChild *TreeNode
	parent     *TreeNode
	str        string
}

func NewTreeNode(s string) *TreeNode {
	return &TreeNode{str: s}
}

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(node *TreeNode) *BinaryTree {
	return &BinaryTree{root: node}
}

func (bt *BinaryTree) PreOrder(current *TreeNode) {
	if current == nil {
		return
	}
	fmt.Print(current.str, " ")
	bt.PreOrder(current.leftChild)
	bt.PreOrder(current.rightChild)
}

func (bt *BinaryTree) InOrder(current *TreeNode) {
	if current == nil {
		return
	}
	bt.InOrder(current.leftChild)
	fmt.Print(current.str, " ")
	bt.InOrder(current.rightChild)
}

func (bt *BinaryTree) PostOrder(current *TreeNode) {
	if current == nil {
		return
	}
	bt.PostOrder(current.leftChild)
	bt.PostOrder(current.rightChild)
	fmt.Print(current.str, " ")
}

func (bt *BinaryTree) LevelOrder() {
	if bt.root == nil {
		return
	}

	queue := []*TreeNode{bt.root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Print(current.str, " ")

		if current.leftChild != nil {
			queue = append(queue, current.leftChild)
		}

		if current.rightChild != nil {
			queue = append(queue, current.rightChild)
		}
	}
}

func (bt *BinaryTree) leftMost(current *TreeNode) *TreeNode {
	for current.leftChild != nil {
		current = current.leftChild
	}
	return current
}

func (bt *BinaryTree) InOrderSuccessor(current *TreeNode) *TreeNode {
	if current.rightChild != nil {
		return bt.leftMost(current.rightChild)
	}

	successor := current.parent

	for successor != nil && current == successor.rightChild {
		current = successor
		successor = successor.parent
	}

	return successor
}

func (bt *BinaryTree) InOrderByParent(root *TreeNode) {
	current := bt.leftMost(root)

	for current != nil {
		fmt.Print(current.str, " ")
		current = bt.InOrderSuccessor(current)
	}
}

func (bt *BinaryTree) rightMost(current *TreeNode) *TreeNode {
	for current.rightChild != nil {
		current = current.rightChild
	}
	return current
}

func (bt *BinaryTree) InOrderPredecessor(current *TreeNode) *TreeNode {
	if current.leftChild != nil {
		return bt.rightMost(current.leftChild)
	}

	predecessor := current.parent

	for predecessor != nil && current == predecessor.leftChild {
		current = predecessor
		predecessor = predecessor.parent
	}

	return predecessor
}

func (bt *BinaryTree) InOrderReverse(root *TreeNode) {
	current := bt.rightMost(root)

	for current != nil {
		fmt.Print(current.str, " ")
		current = bt.InOrderPredecessor(current)
	}
}
