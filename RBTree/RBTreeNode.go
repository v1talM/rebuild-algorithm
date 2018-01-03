package RBTree

const (
	red bool = true
	black bool = false
)
type RBTreeNode struct {
	data int
	parent,
	lChild,
	rChild *RBTreeNode
	color bool
}

func NewRBTreeNode(data int, color bool) *RBTreeNode {
	return &RBTreeNode{
		data,nil,nil,nil,color,
	}
}

func (this *RBTreeNode) GetData() int {
	return this.data
}

func (this *RBTreeNode) SetData(data int)  {
	this.data = data
}

func (this *RBTreeNode) HasParent() bool {
	return this.GetParent() != nil
}

func (this *RBTreeNode) GetParent() *RBTreeNode {
	return this.parent
}

func (this *RBTreeNode) SetParent(node *RBTreeNode)  {
	this.parent = node
}

func (this *RBTreeNode) HasLChild() bool {
	return this.GetLChild() != nil
}

func (this *RBTreeNode) GetLChild() *RBTreeNode {
	return this.lChild
}

func (this *RBTreeNode) SetLChild(node *RBTreeNode)  {
	this.lChild = node
}

func (this *RBTreeNode) HasRChild() bool {
	return this.GetRChild() != nil
}

func (this *RBTreeNode) GetRChild() *RBTreeNode {
	return this.rChild
}

func (this *RBTreeNode) SetRChild(node *RBTreeNode)  {
	this.rChild = node
}

func (this *RBTreeNode) GetGrandParent() *RBTreeNode {
	if	this.HasParent() {
		return this.GetParent().GetParent()
	}
	return nil
}

func (this *RBTreeNode) GetUncle() *RBTreeNode {
	if	this.HasParent() {
		if this.GetParent() == this.GetGrandParent().GetLChild() {
			return this.GetGrandParent().GetRChild()
		} else {
			return this.GetGrandParent().GetLChild()
		}
	}
	return nil
}

func (this *RBTreeNode) GetSiblings() *RBTreeNode {
	if this.HasParent() {
		if this == this.GetParent().GetLChild() {
			return this.GetParent().GetRChild()
		}
		return this.GetParent().GetLChild()
	}
	return nil
}

func (this *RBTreeNode) Rotation(isLeftRotate bool) *RBTreeNode {
	var root *RBTreeNode
	if this == nil {
		return root
	}
	if !isLeftRotate && !this.HasLChild() {
		return root
	} else if isLeftRotate && !this.HasRChild(){
		return root
	}
	parent := this.GetParent()
	var isleft bool
	if parent != nil {
		isleft = parent.GetLChild() == this
	}
	if isLeftRotate {
		grandSon := this.GetRChild().GetLChild()
		this.GetRChild().SetLChild(this)
		this.SetParent(this.GetRChild())
		this.SetRChild(grandSon)
	} else {
		grandSon := this.GetLChild().GetRChild()
		this.GetLChild().SetRChild(this)
		this.SetParent(this.GetLChild())
		this.SetLChild(grandSon)
	}

	// 判断是否换了根节点
	if parent == nil {
		this.GetParent().parent = nil
		root = this.GetParent()
	} else {
		if isleft {
			parent.SetLChild(this.GetParent())
		} else {
			parent.SetRChild(this.GetParent())
		}
		this.parent.parent = parent
	}
	return root
}