package main

import "fmt"

func (bst *binarySearchTree) preOrderDfsTraversal(root *node) {
	if root == nil {
		return
	}
	fmt.Printf("| %v | ->", root.value)
	bst.preOrderDfsTraversal(root.left)
	bst.preOrderDfsTraversal(root.right)
}
