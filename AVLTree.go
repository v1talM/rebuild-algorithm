package main

import (

	"rebuild-algorithm/AVLTree"
	"fmt"
)

func main() {
	data := []int{10,4,3,7,5,6,9,8}
	avl := AVLTree.NewAVLTree(data...)
	fmt.Println("中序遍历结果:")
	avl.InOrderWalk()
	fmt.Println("\n打印该AVL树:")
	avl.PrintBinaryTree(1)
	fmt.Println("当前AVL树的根结点为:", avl.GetRoot())
}
