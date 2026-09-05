// This code assumed we are not going to store duplicate values (but if case is like that then we could also store the count too using minor tweaking in approach)
package main

import (
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

type binarySearchTree struct {
	head *node
}

func (bst *binarySearchTree) insertNode(n *node) {
	currNode := bst.head
	if currNode == nil {
		bst.head = n
		return
	}
	for currNode != nil {
		if currNode.value > n.value {
			//left

			if currNode.left == nil {
				currNode.left = n
				break
			}
			currNode = currNode.left
		} else if currNode.value < n.value {
			// right
			if currNode.right == nil {
				currNode.right = n
				break
			}
			currNode = currNode.right
		}
	}
}

func (bst *binarySearchTree) deleteNode(root *node, value int) *node {
	if root == nil {
		return nil
	}

	if value > root.value {
		root.right = bst.deleteNode(root.right, value)
	} else if value < root.value {
		root.left = bst.deleteNode(root.left, value)
	} else {
		// Value matched

		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}
		smallestInRightSubtree := smallestNodeInTree(root.right)
		root.value = smallestInRightSubtree.value
		root.right = bst.deleteNode(root.right, smallestInRightSubtree.value)
	}
	return root
}
func smallestNodeInTree(n *node) *node {
	currNode := n
	for currNode.left != nil {
		currNode = currNode.left
	}
	return currNode
}

// This function prints the tree in pre order traversal manner
func (bst *binarySearchTree) printTree(root *node) {

	if root == nil {
		return
	}

	fmt.Println(root.value)
	bst.printTree(root.left)
	bst.printTree(root.right)
}
