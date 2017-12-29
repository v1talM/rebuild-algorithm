package BinaryTree

type BinaryTreeNode struct {
	data int
	parent,
	lChild,
	rChild *BinaryTreeNode
}

func NewBinaryTreeNode(data int) *BinaryTreeNode {
	return &BinaryTreeNode{data, nil, nil, nil}
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
	return this.GetLChild() != nil
}

func (this *BinaryTreeNode) HasRChild() bool {
	return this.GetRChild() != nil
}

func (this *BinaryTreeNode) HasParent() bool {
	return this.parent != nil
}

func (this *BinaryTreeNode) GetParent() *BinaryTreeNode {
	if this.HasParent() {
		return this.parent
	}
	return nil
}

func (this *BinaryTreeNode) SetParent(node *BinaryTreeNode)  {
	this.parent = node
}