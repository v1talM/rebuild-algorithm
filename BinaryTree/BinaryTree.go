package BinaryTree

import (
	"fmt"
)

type BinaryTree struct {
	root *BinaryTreeNode
}

func (this *BinaryTree) GetRoot() *BinaryTreeNode {
	return this.root
}

func (this *BinaryTree) InOrderWalk()  {
	InOrderWalk(this.GetRoot())
}

func (this *BinaryTree) InsertNode(data int)  {
	root := this.root
	insertNode(root, data)
}

func (this *BinaryTree) SearchNode(data int) *BinaryTreeNode {
	target := searchNode(this.GetRoot(), data)
	if target != nil {
		return target
	}
	return nil
}

func (this *BinaryTree) DeleteNode(data int)  {
	target := searchNode(this.GetRoot(), data)
	deleteNode(target)
}

func (this *BinaryTree) PrintBinaryTree(height int)  {
	printBinaryTree(this.GetRoot(), height)
}

func NewBinarySearchTree(data...int) *BinaryTree {
	root := NewBinaryTreeNode(data[0])
	bst := &BinaryTree{root: root}
	for i := 1; i < len(data); i++ {
		bst.InsertNode(data[i])
	}
	return bst
}

func InOrderWalk(node *BinaryTreeNode)  {
	if node != nil {
		InOrderWalk(node.GetLChild())
		fmt.Printf("%v ", node.GetData())
		InOrderWalk(node.GetRChild())
	}
}

func insertNode(node *BinaryTreeNode, data int)  {
	if data < node.GetData() {
		if node.HasLChild() {
			insertNode(node.GetLChild(), data)
		} else {
			newNode := NewBinaryTreeNode(data)
			newNode.SetParent(node)
			node.SetLChild(newNode)
		}
	} else {
		if node.HasRChild() {
			insertNode(node.GetRChild(), data)
		} else {
			newNode := NewBinaryTreeNode(data)
			newNode.SetParent(node)
			node.SetRChild(newNode)
		}
	}
}

func searchNode(node *BinaryTreeNode, data int) *BinaryTreeNode {
	for node != nil && node.GetData() != data {
		if data < node.GetData() {
			node = node.GetLChild()
			return searchNode(node, data)
		} else {
			node = node.GetRChild()
			return searchNode(node, data)
		}
	}
	return node
}

func printBinaryTree(node *BinaryTreeNode, height int)  {
	if node != nil {
		printBinaryTree(node.GetRChild(), height + 1)
		for i := 0; i < height; i++ {
			fmt.Printf("    ")
		}
		fmt.Println(node.GetData())
		printBinaryTree(node.GetLChild(), height + 1)
	}
}

func minNode(node *BinaryTreeNode) *BinaryTreeNode {
	for node.GetLChild() != nil {
		node = node.GetLChild()
	}
	return node
}

func deleteNode(node *BinaryTreeNode)  {
	parent := node.GetParent()
	origin := node
	if node.HasLChild() && node.HasRChild() {
		min := minNode(node.GetRChild())
		node.SetData(min.GetData())
		deleteNode(min)
	} else if node.HasLChild() {
		node = node.GetLChild()
	} else if node.HasRChild() {
		node = node.GetRChild()
	} else {
		node = nil
	}
	if node != nil {
		node.SetParent(parent)
	}
	if parent.GetLChild() == origin {
		parent.SetLChild(node)
	}else {
		parent.SetRChild(node)
	}
}