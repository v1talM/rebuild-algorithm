package BinaryTree

type BinaryTreeNode struct {
	data int
	lChild,
	rChild *BinaryTreeNode
}

func NewBinaryTreeNode(data int) *BinaryTreeNode {
	return &BinaryTreeNode{data, nil, nil}
}

func (this *BinaryTreeNode) GetData() int {
	return this.data
}

func (this *BinaryTreeNode) SetData(data int)  {
	this.data = data
}

func (this *BinaryTreeNode) SetLChild(node *BinaryTreeNode)  {
	this.lChild = node
}

func (this *BinaryTreeNode) SetRChild(node *BinaryTreeNode)  {
	this.rChild = node
}

func (this *BinaryTreeNode) GetLChild() *BinaryTreeNode {
	return this.lChild
}

func (this *BinaryTreeNode) GetRChild() *BinaryTreeNode {
	return this.rChild
}

func (this *BinaryTreeNode) HasLChild() bool {
	if this.GetLChild() != nil {
		return true
	}
	return false
}

func (this *BinaryTreeNode) HasRChild() bool {
	if this.GetRChild() != nil {
		return true
	}
	return false
}