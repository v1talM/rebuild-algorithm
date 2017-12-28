package main

import (
	"rebuild-algorithm/BinaryTree"
	"fmt"
)

func main() {
	fmt.Println("创建一棵二叉查找树，插入[4,5,3,2,6,1]")
	bad_bst := BinaryTree.NewBinarySearchTree(4,5,3,2,6,1)
	fmt.Println("中序遍历结果:")
	bad_bst.InOrderWalk()
	fmt.Println("\n打印该二叉查找树:")
	bad_bst.PrintBinaryTree(1)
}
