package main

import (
	"rebuild-algorithm/BinaryTree"
	"fmt"
)

func main() {
	fmt.Println("创建一棵二叉查找树")
	data := []int{10,4,3,7,5,6,9,8}
	bst := BinaryTree.NewBinarySearchTree(data...)

	fmt.Println("中序遍历结果:")
	bst.InOrderWalk()
	fmt.Println("\n打印该二叉查找树:")
	bst.PrintBinaryTree(1)
}
