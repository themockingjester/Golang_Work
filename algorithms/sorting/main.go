package main

import "fmt"

func main() {
	arr := []int{34, 57, 4, -1, 0, 5, 7, 3, 5, 0, 34, 78}

	// fmt.Printf("Here is your array with selection sort: %+v", selectionSort(arr))

	// fmt.Printf("Here is your array with merge sort: %+v", mergeSort(arr))
	// fmt.Printf("Here is your array with insertion sort: %+v", insertionSort(arr))
	fmt.Printf("Here is your array with quick sort: %+v", quickSort(arr, 0, len(arr)-1))

}
