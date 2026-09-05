package main

import "fmt"

func main() {

	//Intialise a tree
	bst := binarySearchTree{head: nil}

	//Inserting nodes
	bst.insertNode(&node{value: 5})
	bst.insertNode(&node{value: 7})
	bst.insertNode(&node{value: 3})
	bst.insertNode(&node{value: 2})
	bst.insertNode(&node{value: 4})
	bst.insertNode(&node{value: 11})
	bst.insertNode(&node{value: 10})
	bst.insertNode(&node{value: 6})
	//printing our tree

	fmt.Println("\n ------------ Here we are printing tree: --------------")
	bst.printTree(bst.head)

	//bfs traversal

	bst.bfsTraversal(bst.head, []node{})

	// preorder travsersal
	fmt.Println("\n ------------ Here we are preorder dfs traversal: --------------")

	bst.preOrderDfsTraversal(bst.head)

	// inorder travsersal
	fmt.Println("\n ------------ Here we are inorder dfs traversal: --------------")

	bst.inorderDfsTraversal(bst.head)

	// postorder travsersal
	fmt.Println("\n ------------ Here we are postorder dfs traversal: --------------")

	bst.postOrderDfsTraversal(bst.head)

	// removing node

	fmt.Println("\n ------------ Removing value 7 from tree: --------------")

	bst.deleteNode(bst.head, 7)

	// inorder travsersal
	fmt.Println("\n ------------ Here we are inorder dfs traversal: --------------")

	bst.inorderDfsTraversal(bst.head)

}
