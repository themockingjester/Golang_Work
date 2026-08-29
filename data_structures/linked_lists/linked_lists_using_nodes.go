// This is a simple implementation of a Linked List data structure in Go. A linked list is a linear data structure where each element (node) contains a value and a pointer to the next node in the sequence. It allows for efficient insertion and deletion of elements, as well as dynamic memory allocation. This implementation provides basic operations such as adding, searching, deleting, and printing elements in the linked list. we are using nodes as a underlying data structure for creating this linked list.
package main

import (
	"fmt"
)

// node represents a single element in the linked list. It contains a Value of any type and a pointer to the next node in the list.
type node struct {
	Value any
	Next  *node
}

// LinkedList represents a singly linked list data structure. It maintains a pointer to the head node and keeps track of the size of the list.
type LinkedList struct {
	Head *node
	Size int
}

// Function to create a new node with the given value. It initializes the Value field with the provided value and sets the Next pointer to nil, indicating that it is the last node in the list.
func createNode(value any) *node {
	newNode := &node{Value: value, Next: nil}
	return newNode
}

// IntialiseLinkedList creates a new LinkedList and returns a pointer to it. It initializes the head of the list to nil and sets the size to 0.
func IntialiseLinkedList() *LinkedList {
	return &LinkedList{Head: nil, Size: 0}
}

// Add appends a new node with the given value to the end of the linked list. It traverses the list to find the last node and updates its Next pointer to point to the newly created node. If the list is empty, it sets the new node as the head of the list.
func (l *LinkedList) Add(value any) bool {
	node := createNode(value)
	currNode := l.Head

	if currNode == nil {
		l.Head = node
		l.Size++
		return true
	}

	for currNode.Next != nil {
		currNode = currNode.Next
	}
	currNode.Next = node
	l.Size++
	return true
}

// PrintList prints the elements of the linked list in order. It starts from the head node and traverses through each node, printing its value until it reaches the end of the list.
func (l *LinkedList) PrintList() {
	fmt.Printf("------------- Here is your Linked list (with size : %d) \n ", l.Size)
	currNode := l.Head
	for currNode != nil {
		fmt.Printf("%v ->", currNode.Value)
		currNode = currNode.Next
	}

	fmt.Printf("\n End of the list ---------------------\n")
}

// Function to search for a value in the linked list. It traverses the list and checks each node's value against the target value. If found, it returns true along with the index of the node; otherwise, it returns false and -1.
func (l *LinkedList) Peek(value any) (bool, int) {
	currNode := l.Head
	if currNode == nil {
		return false, -1
	}
	index := 0
	for currNode != nil {
		if currNode.Value == value {
			return true, index
		}
		currNode = currNode.Next
		index++
	}
	return false, -1
}

// Function to delete a node from the linked list based on the value. It traverses the list to find the node with the specified value and removes it by updating the next pointer of the previous node. If the node is found and deleted, it returns true; otherwise, it returns false.
func (l *LinkedList) Delete(value any) bool {
	currNode := l.Head
	if currNode == nil {
		return false
	}
	if currNode.Value == value {
		l.Head = currNode.Next
		l.Size--
		return true
	}
	for currNode != nil {
		if currNode.Next == nil {
			return false
		} else if currNode.Next.Value == value {
			currNode.Next = currNode.Next.Next
			l.Size--
			return true
		}
		currNode = currNode.Next
	}
	return false
}
func testLinkedList() {

	// Intialising Linked List

	myList := IntialiseLinkedList()

	//adding items to the linked list

	myList.Add(2)
	myList.Add(3)
	myList.Add(76)
	myList.Add(56)
	myList.Add(34)

	// Searching for an item in the list
	valueToSearch := 34
	isPresent, index := myList.Peek(valueToSearch)
	fmt.Printf("Is %v present in the list? %v\n", valueToSearch, isPresent)
	if isPresent {
		fmt.Printf("Index of %v in the list: %d\n", valueToSearch, index)
	}

	// Deleting an item from the list
	valueToDelete := 56
	isDeleted := myList.Delete(valueToDelete)
	fmt.Printf("Was %v deleted from the list? %v\n", valueToDelete, isDeleted)

	valueToDelete = 100
	isDeleted = myList.Delete(valueToDelete)
	fmt.Printf("Was %v deleted from the list? %v\n", valueToDelete, isDeleted)

	myList.Add(3456)
	myList.PrintList()

}
