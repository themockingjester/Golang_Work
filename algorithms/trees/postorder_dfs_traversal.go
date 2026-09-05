package main

import "fmt"

func (bst *binarySearchTree) postOrderDfsTraversal(root *node) {
	if root == nil {
		return
	}
	bst.postOrderDfsTraversal(root.left)
	bst.postOrderDfsTraversal(root.right)
	
	fmt.Printf("| %v | -> ", root.value)
}
