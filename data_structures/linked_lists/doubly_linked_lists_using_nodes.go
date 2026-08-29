// This is a simple implementation of a Doubly Linked List data structure in Go. A Doubly Linked List is a collection of nodes where each node contains a value and two pointers: one pointing to the previous node and another pointing to the next node. This allows for efficient insertion, deletion, and traversal of elements in both directions (forward and backward). The implementation includes methods for adding, deleting, searching, and printing elements in the list.
package main

import "fmt"

// doublyNode represents a single element in the doubly linked list. It contains a Value of any type, a pointer to the previous node, and a pointer to the next node in the list.
type doublyNode struct {
	Value any
	Prev  *doublyNode
	Next  *doublyNode
}

// DoublyLinkedList represents a doubly linked list data structure. It maintains pointers to the head and tail nodes and keeps track of the size of the list.
type DoublyLinkedList struct {
	Head *doublyNode
	Tail *doublyNode
	Size int
}

// Function to create a new doubly node with the given value. It initializes the Value field with the provided value and sets the Prev and Next pointers to nil, indicating that it is the last node in the list.
func createDoublyNode(value any) *doublyNode {
	newNode := &doublyNode{Value: value, Prev: nil, Next: nil}
	return newNode
}

// IntialiseDoublyLinkedList creates a new DoublyLinkedList and returns a pointer to it. It initializes the head and tail of the list to nil and sets the size to 0.
func IntialiseDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{Head: nil, Tail: nil, Size: 0}
}

//Function to add a new node with the given value to the end of the doubly linked list. It traverses the list to find the last node and updates its Next pointer to point to the newly created node. It also updates the Prev pointer of the new node to point back to the previous last node. If the list is empty, it sets the new node as both the head and tail of the list.
func (dll *DoublyLinkedList) Add(value any) bool {
	node := createDoublyNode(value)
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		dll.Size++
		return true
	}

	currNode := dll.Head
	for currNode.Next != nil {
		currNode = currNode.Next
	}
	currNode.Next = node
	node.Prev = currNode
	dll.Tail = node
	dll.Size++
	return true
}

//Function to print the elements of the doubly linked list in order. It starts from the head node and traverses through each node, printing its value until it reaches the end of the list.
func (dll *DoublyLinkedList) PrintList() {
	fmt.Printf(" ------------ Here is your Doubly Linked list (with size : %d) \n ", dll.Size)

	currNode := dll.Head
	for currNode != nil {
		fmt.Printf("%v ->", currNode.Value)
		currNode = currNode.Next
	}
	fmt.Println("\n End of the list ---------------------")
}

// Function to search for a value in the doubly linked list. It traverses the list from the head node, comparing each node's value with the target value. If a match is found, it returns true along with the index of the node; otherwise, it returns false and -1.
func (dll *DoublyLinkedList) Peek(value any) (bool, int) {
	currNode := dll.Head
	index := 0
	for currNode != nil {
		if currNode.Value == value {
			return true, index
		}
		index++
		currNode = currNode.Next
	}
	return false, -1
}

// Function to delete a node with the specified value from the doubly linked list. It traverses the list, searching for the node with the matching value. If found, it updates the Next pointer of the previous node and the Prev pointer of the next node to remove the target node from the list. It also handles special cases for deleting the head or tail nodes. The function returns true if the deletion was successful; otherwise, it returns false.
func (dll *DoublyLinkedList) Delete(value any) bool {
	currNode := dll.Head
	if currNode == nil {
		return false
	} else if dll.Head == dll.Tail && currNode.Value == value {
		dll.Head = nil
		dll.Tail = nil
		dll.Size--
		return true

	}

	for currNode.Next != nil {
		if currNode.Next.Value == value {
			if currNode.Next == dll.Tail {
				dll.Tail = currNode
				currNode.Next = nil
				dll.Size--
				return true
			}

			currNode.Next = currNode.Next.Next
			currNode.Next.Next.Prev = currNode
			dll.Size--
			return true
		}
		currNode = currNode.Next
	}

	return false
}

// testDoublyLinkedList is a test function that demonstrates the functionality of the DoublyLinkedList data structure. It initializes a new doubly linked list, adds several values to it, searches for a specific value, deletes a value, and prints the current state of the list after each operation.
func testDoublyLinkedList() {

	// Initialise a new doubly linked list

	myDoublyLinkedList := IntialiseDoublyLinkedList()

	// Adding items to the doubly linked list
	myDoublyLinkedList.Add(2)
	myDoublyLinkedList.Add(3)
	myDoublyLinkedList.Add(76)
	myDoublyLinkedList.Add(100)
	myDoublyLinkedList.Add(200)
	myDoublyLinkedList.Add(2)

	// Searching for a value in the doubly linked list
	valueToSearch := 100
	found, index := myDoublyLinkedList.Peek(valueToSearch)
	if found {
		fmt.Printf("Value %v found at index %d\n", valueToSearch, index)
	} else {
		fmt.Printf("Value %v not found in the list\n", valueToSearch)
	}

	// Deleting a value from the doubly linked list
	valueToDelete := 76
	deleted := myDoublyLinkedList.Delete(valueToDelete)
	if deleted {
		fmt.Printf("Value %v deleted from the list\n", valueToDelete)
	} else {
		fmt.Printf("Value %v not found in the list\n", valueToDelete)
	}

	// Print the current state of the doubly linked list
	myDoublyLinkedList.PrintList()

}
