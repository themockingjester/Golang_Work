// Here we have implemented FIFO queue using nodes
package main

import "fmt"

type node struct {
	value any
	next  *node
}

type queueUsingNode struct {
	front *node
	back  *node
	size  int
}

func createNode(value any) *node {
	return &node{
		value: value,
		next:  nil,
	}
}
func intialiseQueueUsingNode() *queueUsingNode {
	return &queueUsingNode{front: nil, back: nil, size: 0}
}
func (q *queueUsingNode) enqueue(value any) bool {
	newNode := createNode(value)
	if q.size == 0 {

		q.front = newNode

	} else {
		q.back.next = newNode

	}
	q.size++
	q.back = newNode
	return true
}

func (q *queueUsingNode) dequeue() (bool, any) {
	if q.size == 0 {
		return false, nil
	} else if q.front == q.back {
		valueToReturn := q.front.value
		q.front = nil
		q.back = nil
		q.size--
		return true, valueToReturn
	}
	valueToReturn := q.front.value
	q.front = q.front.next
	q.size--
	return true, valueToReturn
}

func (q *queueUsingNode) Search(value any) (bool, int) {
	if q.size == 0 {
		return false, -1
	}
	index := 0
	currNode := q.front
	for currNode != nil {
		if currNode.value == value {
			return true, index
		}
		currNode = currNode.next
		index++
	}
	return false, -1
}

func (q *queueUsingNode) PrintQueue() {
	fmt.Println("---------------- Here is your FIFO Queue --------------")
	if q.size == 0 {
		return
	}
	fmt.Printf("Front: %v, back: %v\n", q.front.value, q.back.value)
	currNode := q.front

	for currNode != nil {
		fmt.Printf("%v<-", currNode.value)
		currNode = currNode.next
	}

	fmt.Printf("\n -------------------- End of the List ----------------\n")
}

func testQueueUsingNodes() {
	// Intialising Queue

	myQueue := intialiseQueueUsingNode()

	// Adding items

	myQueue.enqueue(453)
	myQueue.enqueue(2)
	myQueue.enqueue(78)
	myQueue.enqueue(234)
	myQueue.enqueue(577)

	// Searching
	valueToSearch := 234
	ifFound, foundIndex := myQueue.Search(valueToSearch)
	if ifFound {
		fmt.Printf("Searched Value: %v, found at index: %v\n", valueToSearch, foundIndex)
	} else {
		fmt.Printf("Searched Value: %v, not found \n", valueToSearch)

	}

	// Dequeing
	removedValueSuccess, valueRemoved := myQueue.dequeue()
	if removedValueSuccess {
		fmt.Printf("removed value: %v", valueRemoved)
	} else {
		fmt.Printf("nothing to remove")
	}

	// Printing Queue
	myQueue.PrintQueue()
}
