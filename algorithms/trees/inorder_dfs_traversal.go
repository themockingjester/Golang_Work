package main

import "fmt"

func (bst *binarySearchTree) inorderDfsTraversal(root *node) {
	if root == nil {
		return
	}
	bst.inorderDfsTraversal(root.left)
	fmt.Printf("| %v | -> ", root.value)
	bst.inorderDfsTraversal(root.right)
}
