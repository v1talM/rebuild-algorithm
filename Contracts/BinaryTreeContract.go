package Contracts

import (
	"rebuild-algorithm/BinaryTree"
	"rebuild-algorithm/AVLTree"
)

type BaseTreeNodeContract interface{
	GetData() int
	SetData(data int)
	HasLChild() bool
	HasRChild() bool
	HasParent() bool
}

type BaseTreeContract interface{
	//中序遍历二叉树
	InOrderWalk()
	//插入新结点
	InsertNode(data int)
	//删除结点
	DeleteNode(data int)
	//树状打印二叉树
	PrintBinaryTree(height int)
}

type BinaryTreeNodeContract interface{
	BaseTreeNodeContract
	GetLChild() *BinaryTree.BinaryTreeNode
	SetLChild(node *BinaryTree.BinaryTreeNode)
	GetRChild() *BinaryTree.BinaryTreeNode
	SetRChild(node *BinaryTree.BinaryTreeNode)
	GetParent() *BinaryTree.BinaryTreeNode
	SetParent(node *BinaryTree.BinaryTreeNode)
}

type AVLTreeNodeContract interface {
	BaseTreeNodeContract
	GetLChild() *AVLTree.AVLTreeNode
	SetLChild(node *AVLTree.AVLTreeNode)
	SetRChild(node *AVLTree.AVLTreeNode)
	GetRChild() *AVLTree.AVLTreeNode
	SetParent(node *AVLTree.AVLTreeNode)
	GetParent()*AVLTree.AVLTreeNode
}

type BinaryTreeContract interface {
	BaseTreeContract
	//获取根结点信息
	GetRoot() *BinaryTree.BinaryTreeNode
	//搜索结点
	SearchNode(data int) *BinaryTree.BinaryTreeNode
}

type AVLTreeContract interface {
	BaseTreeContract
	GetRoot() *AVLTree.AVLTreeNode
}
