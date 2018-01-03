package AVLTree

import (
	"fmt"
)

type AVLTree struct{
	root *AVLTreeNode
}

func NewAVLTree(data...int) *AVLTree {
	root := NewAVLTreeNode(data[0])
	tree := &AVLTree{root: root}
	for i := 1; i < len(data); i++ {
		tree.InsertNode(data[i])
	}
	return tree
}

func (this *AVLTree) GetRoot() *AVLTreeNode {
	return this.root
}

func (this *AVLTree) InOrderWalk()  {
	InOrderWalk(this.GetRoot())
}

func (this *AVLTree) InsertNode(data int)  {
	this.root = insertNode(this.root, data)
}

func (this *AVLTree) DeleteNode(data int)  {
	target := searchNode(this.GetRoot(), data)
	deleteNode(target)
}

func (this *AVLTree) PrintBinaryTree(height int)  {
	printBinaryTree(this.GetRoot(), 1)
}

func InOrderWalk(node *AVLTreeNode)  {
	if node != nil {
		InOrderWalk(node.GetLChild())
		fmt.Printf("%v ", node.GetData())
		InOrderWalk(node.GetRChild())
	}
}

func printBinaryTree(node *AVLTreeNode, height int)  {
	if node != nil {
		printBinaryTree(node.GetRChild(), height + 1)
		for i := 0; i < height; i++ {
			fmt.Printf("    ")
		}
		fmt.Println(node.GetData())
		printBinaryTree(node.GetLChild(), height + 1)
	}
}

func searchNode(node *AVLTreeNode, data int) *AVLTreeNode {
	if node != nil && node.GetData() != data {
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

func insertNode(node *AVLTreeNode, data int) *AVLTreeNode {
	if node == nil {
		return NewAVLTreeNode(data)
	} else if data < node.GetData() {
		node.SetLChild(insertNode(node.GetLChild(), data))
	} else {
		node.SetRChild(insertNode(node.GetRChild(), data))
	}

	if height(node.GetLChild()) - height(node.GetRChild()) == 2 {
		if data < node.GetLChild().GetData() {
			node = rightRotation(node)
		} else {
			node = leftRightRotation(node)
		}
	}
	if height(node.GetRChild()) - height(node.GetLChild()) == 2 {
		if data > node.GetRChild().GetData() {
			node = leftRotation(node)
		} else {
			node = rightLeftRotation(node)
		}
	}

	node.height = max(height(node.GetLChild()), height(node.GetRChild())) + 1
	return node
}

func deleteNode(node *AVLTreeNode) {
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
	}else {
		node = nil
	}

	if node != nil {
		node.SetParent(parent)
	}
	if parent.GetLChild() == origin {
		parent.SetLChild(node)
	} else {
		parent.SetRChild(node)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func leftRotation(x *AVLTreeNode) *AVLTreeNode {
	//x ->node
	//  y
	//    z
	y := x.GetRChild()
	if y.HasLChild() {
		y.GetLChild().SetParent(x)
	}
	x.SetRChild(y.GetLChild())
	y.SetLChild(x)
	x.SetParent(y)
	y.GetRChild().SetParent(y)
	return y
}

func rightRotation(x *AVLTreeNode) *AVLTreeNode {
	//    x ->node
	//  y
	//z
	y := x.GetLChild()
	if y.HasRChild() {
		y.GetRChild().SetParent(x)
	}
	x.SetLChild(y.GetRChild())
	y.SetRChild(x)
	x.SetParent(y)
	y.GetLChild().SetParent(y)
	return y
}

func leftRightRotation(x *AVLTreeNode) *AVLTreeNode {
	//   x
	// y
	//   z
	y := x.GetLChild()
	x.SetLChild(leftRotation(y))
	x = rightRotation(x)
	return x
}

func rightLeftRotation(x *AVLTreeNode) *AVLTreeNode  {
	// x
	//   y
	// z
	y := x.GetRChild()
	x.SetRChild(rightRotation(y))
	x = leftRotation(x)
	return x
}

func height(node *AVLTreeNode) int {
	if node == nil {
		return -1
	}
	lh, rh := 0, 0
	lchild, rchild := node.GetLChild(), node.GetRChild()
	for lchild != nil {
		lh++
		lchild = lchild.GetLChild()
	}
	for rchild != nil {
		rh++
		rchild = rchild.GetRChild()
	}
	return max(lh, rh)
}

func minNode(node *AVLTreeNode) *AVLTreeNode {
	for node.HasLChild() {
		node = node.GetLChild()
	}
	return node
}