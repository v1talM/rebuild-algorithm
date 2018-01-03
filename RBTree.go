package main

import (
	"rebuild-algorithm/RBTree"
	"fmt"
)

func main() {
	rbtree := RBTree.NewRBTree(1)
	fmt.Println("顺序插入：1 2 3 4 5 6 7 8")
	rbtree.InsertNode(rbtree.GetRoot(), 2)
	rbtree.InsertNode(rbtree.GetRoot(), 3)
	rbtree.InsertNode(rbtree.GetRoot(), 4)
	rbtree.InsertNode(rbtree.GetRoot(), 5)
	rbtree.InsertNode(rbtree.GetRoot(), 6)
	rbtree.InsertNode(rbtree.GetRoot(), 7)
	rbtree.InsertNode(rbtree.GetRoot(), 8)
	fmt.Println("\n打印该红黑树:")
	rbtree.PrintBinaryTree(1)
}
