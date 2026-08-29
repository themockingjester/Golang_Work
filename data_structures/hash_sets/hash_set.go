// This is a simple implementation of a HashSet data structure in Go. A HashSet is a collection that stores unique keys and their associated values. It uses a hash function to calculate the index for each key, allowing for efficient insertion, retrieval, and deletion of elements. (Here we used a Array as a underlying data structure for creating this hashset)

package main

import (
	"fmt"
)

// Allowed type for Keys
type KeyType interface {
	string | int
}

// These HashSets can use any type of value
type ValueType interface {
	any
}

// Indivial pair of key and value stored in the HashSet
type Element[k KeyType] struct {
	key   k
	value any
}

// HashSet is a data structure that stores unique keys and their associated values. It uses a hash function to calculate the index for each key, allowing for efficient insertion, retrieval, and deletion of elements.
type HashSet[k KeyType] struct {
	Memory []ValueType
}

// IntialiseHashSet creates a new HashSet with the specified capacity. It initializes the underlying memory for storing elements and returns the created HashSet.
func IntialiseHashSet[k KeyType](capacity int) HashSet[k] {
	memory := make([]ValueType, capacity)
	hashMap := HashSet[k]{Memory: memory}
	return hashMap
}

// CalculateHash generates a final int , this function can be enhanced based on needs
func (h HashSet[k]) calculateHash(key k) int {
	keyAsString := fmt.Sprintf("%s", key)
	result := 0

	for i := 0; i < len(keyAsString); i++ {
		result = result*i + int(keyAsString[i])
	}
	return result % len(h.Memory)
}

// Add inserts a new key-value pair into the HashSet. It calculates the hash for the given key and stores the element in the appropriate index of the underlying memory. If there are already elements at that index, it appends the new element to the existing slice.
func (h HashSet[k]) Add(key k, value ValueType) {
	calculatedHash := h.calculateHash(key)
	fmt.Printf("Calculated Hash for Key %v: %v\n", key, calculatedHash)
	newElement := Element[k]{key: key, value: value}

	if h.Memory[calculatedHash] == nil {
		h.Memory[calculatedHash] = []Element[k]{newElement}

	} else {
		h.Memory[calculatedHash] = append(h.Memory[calculatedHash].([]Element[k]), newElement)
	}
}

// Get retrieves the value associated with the given key from the HashSet. It calculates the hash for the key and searches for the corresponding element in the underlying memory. If found, it returns the value and a boolean indicating success; otherwise, it returns nil and false.
func (h HashSet[k]) Get(key k) (ValueType, bool) {
	calculatedHash := h.calculateHash(key)
	elementsAtHash := h.Memory[calculatedHash]
	for _, element := range elementsAtHash.([]Element[k]) {
		if element.key == key {
			return element.value, true
		}
	}
	return nil, false
}

// Delete removes the key-value pair associated with the given key from the HashSet. It calculates the hash for the key and searches for the corresponding element in the underlying memory. If found, it removes the element and returns true; otherwise, it returns false.
func (h HashSet[k]) Delete(key k) bool {
	calculatedHash := h.calculateHash(key)
	elementsAtHash := h.Memory[calculatedHash]
	newElements := []Element[k]{}
	dataFound := false
	for i, element := range elementsAtHash.([]Element[k]) {
		if element.key != key {
			newElements = append(newElements, elementsAtHash.([]Element[k])[i])

		} else {
			dataFound = true
		}
	}
	h.Memory[calculatedHash] = newElements
	return dataFound
}

func main() {

	// Intialising HashSet with the capacity of 10
	hashSet := IntialiseHashSet[string](10)
	fmt.Printf("%+v", hashSet)

	// Adding data to our set
	hashSet.Add("key1", "value1")
	hashSet.Add("key2", "value2")
	hashSet.Add("key3", "value3")
	hashSet.Add("key4", "value4")
	hashSet.Add("key5", "value5")
	hashSet.Add("key6", "value6")
	hashSet.Add("key7", "value7")
	hashSet.Add("key8", "value8")
	hashSet.Add("key9", "value9")
	hashSet.Add("key10", "value10")
	hashSet.Add("key11", "value11")

	valueFound, _ := hashSet.Get("key1")
	fmt.Printf("Value at Key1: %v\n", valueFound)

	// deleting data from our set
	fmt.Printf("Deleting Key1: %v\n", hashSet.Delete("key1"))
	fmt.Printf("%+v", hashSet)

}
