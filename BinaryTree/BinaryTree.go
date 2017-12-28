package BinaryTree

import "fmt"

type BinaryTree struct {
	root *BinaryTreeNode
}

func (this *BinaryTree) GetRoot() *BinaryTreeNode {
	return this.root
}

func (this *BinaryTree) InOrderWalk()  {
	inOrderWalk(this.root)
}

func (this *BinaryTree) InsertNode(data int)  {
	root := this.root
	insertNode(root, data)
}

func (this *BinaryTree) SearchNode(data int) bool {
	target := searchNode(this.GetRoot(), data)
	if target != nil {
		return true
	}
	return false
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

func inOrderWalk(node *BinaryTreeNode)  {
	if node != nil {
		inOrderWalk(node.GetLChild())
		fmt.Printf("%v ", node.GetData())
		inOrderWalk(node.GetRChild())
	}
}

func insertNode(node *BinaryTreeNode, data int)  {
	if data < node.GetData() {
		if node.HasLChild() {
			insertNode(node.GetLChild(), data)
		} else {
			newNode := NewBinaryTreeNode(data)
			node.SetLChild(newNode)
		}
	} else {
		if node.HasRChild() {
			insertNode(node.GetRChild(), data)
		} else {
			newNode := NewBinaryTreeNode(data)
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