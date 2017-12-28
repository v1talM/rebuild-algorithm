package BinaryTree

type AVLTreeNode struct{
	BinaryTreeNode
	height int
}

func (this *AVLTreeNode) SetHeight(height int)  {
	this.height = height
}

func (this *AVLTreeNode) GetHeight() int {
	return this.height
}

