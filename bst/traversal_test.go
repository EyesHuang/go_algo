package bst

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_PreOrder(t *testing.T) {
	tree := genBST()
	tree.PreOrder(tree.root)
	fmt.Println()
}

func TestBinaryTree_InOrder(t *testing.T) {
	tree := genBST()
	tree.InOrder(tree.root)
	fmt.Println()
}

func TestBinaryTree_PostOrder(t *testing.T) {
	tree := genBST()
	tree.PostOrder(tree.root)
	fmt.Println()
}

func TestBinaryTree_LevelOrder(t *testing.T) {
	tree := genBST()
	tree.LevelOrder()
	fmt.Println()
}

func TestBinaryTree_InOrderByParent(t *testing.T) {
	tree := genBSTWithParent()
	tree.InOrderByParent(tree.root)
	fmt.Println()
}

func TestBinaryTree_InOrderPredecessor(t *testing.T) {
	tree := genBSTWithParent()
	tree.InOrderReverse(tree.root)
	fmt.Println()
}

func genBST() *BinaryTree {
	nodeA := NewTreeNode("A")
	nodeB := NewTreeNode("B")
	nodeC := NewTreeNode("C")
	nodeD := NewTreeNode("D")
	nodeE := NewTreeNode("E")
	nodeF := NewTreeNode("F")
	nodeG := NewTreeNode("G")
	nodeH := NewTreeNode("H")
	nodeI := NewTreeNode("I")

	// construct the Binary Tree
	nodeA.leftChild = nodeB
	nodeA.rightChild = nodeC
	nodeB.leftChild = nodeD
	nodeB.rightChild = nodeE
	nodeE.leftChild = nodeG
	nodeE.rightChild = nodeH
	nodeC.leftChild = nodeF
	nodeF.rightChild = nodeI

	tree := NewBinaryTree(nodeA)
	return tree
}

func genBSTWithParent() *BinaryTree {
	nodeA := NewTreeNode("A")
	nodeB := NewTreeNode("B")
	nodeC := NewTreeNode("C")
	nodeD := NewTreeNode("D")
	nodeE := NewTreeNode("E")
	nodeF := NewTreeNode("F")
	nodeG := NewTreeNode("G")
	nodeH := NewTreeNode("H")
	nodeI := NewTreeNode("I")

	// construct the Binary Tree
	nodeA.leftChild = nodeB
	nodeA.rightChild = nodeC

	nodeB.leftChild = nodeD
	nodeB.rightChild = nodeE

	nodeE.leftChild = nodeG
	nodeE.rightChild = nodeH

	nodeC.leftChild = nodeF
	nodeF.rightChild = nodeI

	nodeB.parent = nodeA
	nodeC.parent = nodeA
	nodeD.parent = nodeB
	nodeE.parent = nodeB
	nodeF.parent = nodeC
	nodeG.parent = nodeE
	nodeH.parent = nodeE
	nodeI.parent = nodeF

	tree := NewBinaryTree(nodeA)
	return tree
}
