// This approach uses a underlying slice (internally uses array) for implementing stack
package main

import "fmt"

type arrayStack struct {
	memory []any
	size   int
	top    int
}

func intialiseArrayStack() *arrayStack {
	return &arrayStack{memory: make([]any, 0), size: 0, top: -1}
}

func (a *arrayStack) Push(value any) bool {
	a.memory = append(a.memory, value)
	a.top = a.size + 1
	a.size++
	return true
}

func (a *arrayStack) Pop() bool {
	if a.size == 0 {
		return false
	}
	a.size--
	a.memory = a.memory[:a.size]
	a.top--
	return true

}

func (a *arrayStack) PrintStack() {
	fmt.Printf("------------ Here s your stack --------------\n")
	if a.size == 0 {
		return
	}
	fmt.Printf("Top: %v, Size: %v\n", a.top, a.size)
	for _, value := range a.memory {
		fmt.Printf("%v->", value)
	}
	fmt.Println("\n------------- End of stack")
}

func (a *arrayStack) Search(valueSearched any) (bool, int) {
	for i, value := range a.memory {
		if value == valueSearched {
			return true, i
		}
	}
	return false, -1
}

func testArrayStack() {
	// Initialsing Stack

	myStack := intialiseArrayStack()

	// Adding data

	myStack.Push(34)
	myStack.Push(67)
	myStack.Push(23)
	myStack.Push(90)

	// Searching Element
	valueToSearch := 785
	found, foundIndex := myStack.Search(valueToSearch)
	if found {
		fmt.Printf("Searched Value; %v, found at index: %v\n", valueToSearch, foundIndex)
	} else {
		fmt.Printf("Searched Value; %v, not found\n", valueToSearch)

	}

	// Popping Element

	myStack.Pop()

	// Printing Stack

	myStack.PrintStack()
}
