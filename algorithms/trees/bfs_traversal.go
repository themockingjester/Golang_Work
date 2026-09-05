package main

import "fmt"

func (bst *binarySearchTree) bfsTraversal(root *node, queue []node) {
	fmt.Println("------------ Here we have BFS traversal:--------")
	currNode := root
	if currNode == nil {
		return
	}
	queue = append(queue, *currNode)
	for len(queue) != 0 {
		poppedVal := queue[0]
		queue = queue[1:]
		fmt.Printf("| %v | ", poppedVal.value)
		if poppedVal.left != nil {
			queue = append(queue, *poppedVal.left)
		}
		if poppedVal.right != nil {
			queue = append(queue, *poppedVal.right)
		}
	}

}
