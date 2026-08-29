// This approach uses nodes for implementation of stack
package main

import "fmt"

type node struct {
	Value any
	Next  *node
}

func createNode(value any) *node {
	newNode := &node{Value: value, Next: nil}
	return newNode
}

type nodeStack struct {
	Top    *node
	Bottom *node
	Size   int
}

// Function to intialise the stack
func intialiseStack() *nodeStack {
	return &nodeStack{Top: nil, Size: 0}
}

// Function to pushing a elment in the stack
func (s *nodeStack) Push(value any) bool {
	newNode := createNode(value)
	if s.Size == 0 {
		s.Top = newNode
		s.Bottom = newNode
		s.Size++
		return true
	}
	s.Top.Next = newNode
	s.Top = newNode
	s.Size++
	return true
}

// Function to Popping the element
func (s *nodeStack) Pop() bool {
	if s.Size == 0 {
		return false
	} else if s.Top == s.Bottom {
		s.Top = nil
		s.Bottom = nil
		s.Size -= 1
		return true
	}
	currNode := s.Bottom
	for currNode.Next.Next != nil {
		currNode = currNode.Next
	}
	currNode.Next = nil
	s.Top = currNode
	s.Size--
	return true
}

func (s *nodeStack) PrintStack() {
	fmt.Printf("--------- Here is you stack ----- \n")
	if s.Size == 0 {
		return
	}
	fmt.Printf("Top: %v, Bottom: %v\n", s.Top.Value, s.Bottom.Value)
	currNode := s.Bottom
	for currNode != nil {
		fmt.Printf("%v->", currNode.Value)
		currNode = currNode.Next
	}
}

func (s *nodeStack) Search(value any) (bool, int) {
	index := -1
	if s.Size == 0 {
		return false, -1
	}
	index++
	currNode := s.Bottom
	for currNode != nil {
		if currNode.Value == value {
			return true, index
		}
		currNode = currNode.Next
		index++
	}
	return false, -1
}

func testStackUsingNodes() {
	// Initialise stack
	myStack := intialiseStack()

	// // Adding items
	myStack.Push(12)
	myStack.Push(13)
	myStack.Push(2)
	myStack.Push(56)
	myStack.Push(674)

	// Searching value
	valueToSearch := 78554
	isFound, foundIndex := myStack.Search(valueToSearch)
	if isFound {
		fmt.Printf("Searched value: %v, found at index: %v\n", valueToSearch, foundIndex)
	} else {
		fmt.Printf("Searched value: %v, not found \n", valueToSearch)

	}

	// Pop Element

	myStack.Pop()

	// Printing Stack
	myStack.PrintStack()
}
