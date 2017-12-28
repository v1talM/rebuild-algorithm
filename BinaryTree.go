package main

import (
	"rebuild/BinaryTree"
	"fmt"
)

func main() {
	bst := BinaryTree.NewBinarySearchTree(1,2,3,4,5,6,7)
	bst.InOrderWalk()
	fmt.Println("\n二叉树打印：")
	bst.PrintBinaryTree(1)
	fmt.Println("寻找 8 结点：", bst.SearchNode(8))
}
