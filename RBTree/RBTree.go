package RBTree

import "fmt"

const (
	LeftRotation bool = true
	RightRotation bool = false
)

type RBTree struct {
	root *RBTreeNode
}

func NewRBTree(data... int) *RBTree {
	root := NewRBTreeNode(data[0], black)
	for i := 1; i < len(data); i++ {
		insertRBNode(root, data[i])
	}
	return &RBTree{root: root}
}

func (this *RBTree) GetRoot() *RBTreeNode {
	return this.root
}

func (this *RBTree) InOrderWalk()  {
	inOrderWalk(this.root)
}

func (this *RBTree) InsertNode(pnode *RBTreeNode, data int)  {
	if data < pnode.GetData() {
		if pnode.HasLChild() {
			this.InsertNode(pnode.GetLChild(), data)
		} else {
			tmpnode := NewRBTreeNode(data, red)
			tmpnode.SetParent(pnode)
			pnode.SetLChild(tmpnode)
			this.insertCheck(tmpnode)
		}
	} else {
		if pnode.HasRChild() {
			this.InsertNode(pnode.GetRChild(), data)
		} else {
			tmpnode := NewRBTreeNode(data, red)
			tmpnode.SetParent(pnode)
			pnode.SetRChild(tmpnode)
			this.insertCheck(tmpnode)
		}
	}
}

func (this *RBTree) DeleteNode(data int)  {

}

func (this *RBTree) PrintBinaryTree(height int) {
	printBinaryTree(this.GetRoot(), 1)
}

func (this *RBTree) insertCheck(node *RBTreeNode)  {
	if node.GetParent() == nil {
		this.root = node
		this.root.color = black
	}
	if node.HasParent() && node.GetParent().color == red {
		if node.GetUncle() != nil && node.GetUncle().color == red {
			//叔叔结点为红色的情况
			node.GetParent().color = black
			node.GetUncle().color = black
			node.GetGrandParent().color = red
			this.insertCheck(node.GetGrandParent())
		} else {
			//叔叔结点为黑色或者不存在的情况
			isLeft := node == node.GetParent().GetLChild()
			isParentLeft := node.GetParent() == node.GetGrandParent().GetLChild()
			if isLeft && isParentLeft {
				//当前结点以及当前结点的父节点都是左子结点的情况
				node.GetParent().color = black
				node.GetGrandParent().color = red
				this.RightRotation(node.GetGrandParent())
			} else if !isLeft && isParentLeft {
				//父结点为左子结点，当前结点为右子结点
				this.LeftRotation(node.GetParent())
				this.RightRotation(node.GetGrandParent())
				node.color = black
				node.GetLChild().color = red
				node.GetRChild().color = red
			} else if isLeft && !isParentLeft {
				this.RightRotation(node.GetParent())
				this.LeftRotation(node.GetGrandParent())
				node.color = black
				node.GetLChild().color = red
				node.GetRChild().color = red
			} else if !isLeft && !isParentLeft {
				node.GetParent().color = black
				node.GetGrandParent().color = red
				this.LeftRotation(node.GetGrandParent())
			}
		}
	}
}

func (this *RBTree) LeftRotation(node *RBTreeNode)  {
	if tmproot := node.Rotation(LeftRotation); tmproot != nil {
		this.root = tmproot
	}
}

func (this *RBTree) RightRotation(node *RBTreeNode)  {
	if tmproot := node.Rotation(RightRotation); tmproot != nil {
		this.root = tmproot
	}
}

func inOrderWalk(node *RBTreeNode)  {
	if node != nil {
		inOrderWalk(node.GetLChild())
		fmt.Printf("%v  ", node.GetData())
		inOrderWalk(node.GetRChild())
	}
}

func insertRBNode(root *RBTreeNode, data int)  {
	x := root
	var y *RBTreeNode
	z := NewRBTreeNode(data, red)
	for x != nil {
		y = x
		if data < x.GetData() {
			x = x.GetLChild()
		} else {
			x = x.GetRChild()
		}
	}
	z.SetParent(y)
	if z.GetData() < y.GetData() {
		y.SetLChild(z)
	}else {
		y.SetRChild(z)
	}

}

func printBinaryTree(node *RBTreeNode, height int)  {
	if node != nil {
		printBinaryTree(node.GetRChild(), height + 1)
		for i := 0; i < height; i++ {
			fmt.Printf("    ")
		}
		if node.color == red {
			fmt.Printf("%c[1;40;31m%v%c[0m\n\n", 0x1B, node.GetData(), 0x1B)
		} else {
			fmt.Printf("%c[1;30;47m%v%c[0m\n\n", 0x1B, node.GetData(), 0x1B)
		}
		printBinaryTree(node.GetLChild(), height + 1)
	}
}

