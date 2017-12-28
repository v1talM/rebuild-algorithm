package Contracts

import (
	"rebuild-algorithm/BinaryTree"
)

type BinaryTreeNodeContract interface{
	GetData() int
	SetData(data int)
	GetLChild() *BinaryTree.BinaryTreeNode
	SetLChild(node *BinaryTree.BinaryTreeNode)
	GetRChild() *BinaryTree.BinaryTreeNode
	SetRChild(node *BinaryTree.BinaryTreeNode)
	HasLChild() bool
	HasRChild() bool
}

type BinaryTreeContract interface {
	//中序遍历二叉树
	InOrderWalk()
	//获取根结点信息
	GetRoot() *BinaryTree.BinaryTreeNode
	//插入新结点
	InsertNode(data int)
	//搜索结点
	SearchNode(data int) bool
	//树状打印二叉树
	PrintBinaryTree(height int)
}

type AVLTreeContract interface {
	SetHeight(height int)
	GetHeight() int
}

