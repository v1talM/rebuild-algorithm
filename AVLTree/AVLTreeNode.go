package AVLTree

type AVLTreeNode struct{
	data int
	parent,
	lChild,
	rChild *AVLTreeNode
	height int
}

func NewAVLTreeNode(data int) *AVLTreeNode {
	return &AVLTreeNode{
		data,
		nil,
		nil,
		nil,
		0,
	}
}

func (this *AVLTreeNode) GetData() int {
	return this.data
}

func (this *AVLTreeNode) SetData(data int)  {
	this.data = data
}

func (this *AVLTreeNode) HasLChild() bool {
	return this.GetLChild() != nil
}

func (this *AVLTreeNode) HasRChild() bool {
	return this.GetRChild() != nil
}

func (this *AVLTreeNode) HasParent() bool {
	return this.GetParent() != nil
}

func (this *AVLTreeNode) SetParent(node *AVLTreeNode)  {
	this.parent = node
}

func (this *AVLTreeNode) GetParent() *AVLTreeNode {
	if this.parent != nil {
		return this.parent
	}
	return nil
}

func (this *AVLTreeNode) SetLChild(node *AVLTreeNode)  {
	this.lChild = node
}

func (this *AVLTreeNode) GetLChild() *AVLTreeNode {
	if this.lChild != nil {
		return this.lChild
	}
	return nil
}

func (this *AVLTreeNode) SetRChild(node *AVLTreeNode)  {
	this.rChild = node
}

func (this *AVLTreeNode) GetRChild() *AVLTreeNode {
	if this.rChild != nil {
		return this.rChild
	}
	return nil
}


